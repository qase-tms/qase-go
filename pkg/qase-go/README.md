# Qase Go SDK

Qase Go SDK provides comprehensive tools for integrating Go applications with Qase TestOps. The SDK includes domain models, reporters, and testing utilities.

## Features

- **Domain Models**: Complete Go representations of Qase TestOps data structures
- **Reporters**: Flexible reporting system with TestOps and file-based reporters
- **Testing Utilities**: Helper functions for creating test reports
- **Configuration Management**: Easy configuration setup for different environments
- **Error Handling**: Comprehensive error handling and validation

## Quick Start

### Installation

```bash
go get github.com/qase-tms/qase-go/pkg/qase-go
```

### Basic Usage

```go
package main

import (
    "context"
    "testing"
    
    "github.com/qase-tms/qase-go/pkg/qase-go/config"
    "github.com/qase-tms/qase-go/pkg/qase-go/reporters"
    "github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestMyFeature(t *testing.T) {
    // Create configuration
    cfg := &config.Config{
        Mode: "report",
        Report: config.ReportConfig{
            Driver: "local",
            Connection: config.ConnectionConfig{
                Local: config.LocalConfig{
                    Path: "./reports",
                    Format: "json",
                },
            },
        },
    }
    
    // Create reporter
    coreReporter, err := reporters.NewCoreReporter(cfg)
    if err != nil {
        t.Fatal(err)
    }
    
    // Start test run
    ctx := context.Background()
    if err := coreReporter.StartTestRun(ctx); err != nil {
        t.Fatal(err)
    }
    defer coreReporter.CompleteTestRun(ctx)
    
    // Create test result
    result := domain.NewTestResult("My Feature Test")
    result.Title = "My Feature Test"
    result.Fields["description"] = "Testing my feature functionality"
    
    // Create test step
    step := domain.NewTestStep("Perform action")
    step.SetExpectedResult("Action completed successfully")
    step.SetData("input: test data")
    step.SetStatus(domain.StepStatusPassed)
    
    result.AddStep(step)
    
    // Add result to reporter
    if err := coreReporter.AddResult(result); err != nil {
        t.Fatal(err)
    }
}
```

## Components

### Domain Models

The domain package provides Go representations of Qase TestOps data structures:

- `TestResult`: Represents a test execution result
- `TestStep`: Represents a test step with execution details
- `Attachment`: Represents file attachments
- `Status`: Test and step status enums

```go
import "github.com/qase-tms/qase-go/pkg/qase-go/domain"

// Create a test result
result := domain.NewTestResult("My Test")
result.SetRunID(123)
result.Execution.Status = domain.StatusPassed

// Create a test step
step := domain.NewTestStep("My Step")
step.SetExpectedResult("Expected outcome")
step.SetStatus(domain.StepStatusPassed)
```

### Reporters

The reporters package provides flexible reporting capabilities:

- **CoreReporter**: Main reporter with fallback support
- **FileReporter**: File-based reporting
- **TestOpsReporter**: Direct integration with Qase TestOps

```go
import "github.com/qase-tms/qase-go/pkg/qase-go/reporters"

// Create file reporter
fileReporter := reporters.NewFileReporter(reporters.FileReporterConfig{
    Config: cfg,
})

// Add test result
result := domain.NewTestResult("Test")
err := fileReporter.AddResult(result)
```

### Qase Package

The qase package provides high-level testing utilities:

- **Test function**: Simplified test execution with metadata
- **Assert utilities**: Common assertion functions with detailed error capture
- **Reporter integration**: Automatic reporter setup
- **Error Details**: Comprehensive error information capture for failed assertions

```go
import "github.com/qase-tms/qase-go/pkg/qase-go/qase"

// Simple test with metadata
qase.Test(t, qase.TestMetadata{
    DisplayName: "My Test",
    Title: "My Test Title",
    Description: "Test description",
    IDs: []int64{123},
}, func() {
    // Test logic here
    // Results are automatically reported
})
```

#### Enhanced Error Details

The SDK now captures comprehensive error information when assertions fail, including:

- **Error Type**: The type of assertion that failed (e.g., "Should be true", "Not equal")
- **Expected Value**: What was expected
- **Actual Value**: What was actually received
- **Error Message**: Detailed description of the failure
- **File and Line**: Location where the assertion failed

```go
// When an assertion fails, detailed error information is automatically captured
qase.True(t, false, "This will fail")

// The resulting TestResult will contain:
// result.Execution.ErrorDetails.ErrorType = "Should be true"
// result.Execution.ErrorDetails.Expected = "true"
// result.Execution.ErrorDetails.Actual = "false"
// result.Execution.ErrorDetails.File = "/path/to/test.go"
// result.Execution.ErrorDetails.Line = 42
```

#### Available Assertion Functions

All assertion functions automatically capture error details:

```go
// Basic assertions
qase.True(t, value, "message")
qase.False(t, value, "message")
qase.Equal(t, expected, actual, "message")
qase.NotEqual(t, expected, actual, "message")
qase.EqualValues(t, expected, actual, "message")
qase.NotEqualValues(t, expected, actual, "message")

// Error assertions
qase.Error(t, err, "message")
qase.NoError(t, err, "message")
qase.EqualError(t, err, expectedMessage, "message")
qase.ErrorIs(t, err, target, "message")
qase.ErrorAs(t, err, target, "message")

// String assertions
qase.Contains(t, str, substr, "message")
qase.NotContains(t, str, substr, "message")
qase.Regexp(t, rx, str, "message")
qase.NotRegexp(t, rx, str, "message")

// Numeric assertions
qase.Greater(t, e1, e2, "message")
qase.GreaterOrEqual(t, e1, e2, "message")
qase.Less(t, e1, e2, "message")
qase.LessOrEqual(t, e1, e2, "message")

// Collection assertions
qase.Empty(t, object, "message")
qase.NotEmpty(t, object, "message")
qase.Len(t, object, length, "message")
qase.Same(t, expected, actual, "message")
qase.NotSame(t, expected, actual, "message")
qase.Subset(t, list, subset, "message")
qase.NotSubset(t, list, subset, "message")
qase.ElementsMatch(t, listA, listB, "message")
```

## Configuration

The SDK supports flexible configuration through the config package:

```go
import "github.com/qase-tms/qase-go/pkg/qase-go/config"

cfg := &config.Config{
    Mode: "testops", // "report", "testops", or "off"
    Fallback: "report", // Fallback mode if main mode fails
    
    // TestOps configuration
    TestOps: config.TestOpsConfig{
        API: config.APIConfig{
            Token: "your-api-token",
            Host: "api.qase.io",
        },
        Project: "your-project-code",
    },
    
    // Report configuration
    Report: config.ReportConfig{
        Driver: "local",
        Connection: config.ConnectionConfig{
            Local: config.LocalConfig{
                Path: "./reports",
                Format: "json",
            },
        },
    },
}
```

## Examples

### Basic Test with Qase Package

```go
func TestLoginFlow(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        DisplayName: "Login Flow Test",
        Title: "User Login Flow",
        Description: "Test the complete user login process",
        IDs: []int64{456},
        Fields: map[string]string{
            "environment": "staging",
            "browser": "chrome",
        },
    }, func() {
        // Test logic here
        // Results are automatically reported
    })
}
```

### Test with Custom Reporter

```go
func TestWithCustomReporter(t *testing.T) {
    cfg := &config.Config{
        Mode: "report",
        Report: config.ReportConfig{
            Driver: "local",
            Connection: config.ConnectionConfig{
                Local: config.LocalConfig{
                    Path: "./custom-reports",
                    Format: "json",
                },
            },
        },
    }
    
    reporter, err := reporters.NewCoreReporter(cfg)
    if err != nil {
        t.Fatal(err)
    }
    
    ctx := context.Background()
    if err := reporter.StartTestRun(ctx); err != nil {
        t.Fatal(err)
    }
    defer reporter.CompleteTestRun(ctx)
    
    result := domain.NewTestResult("Custom Test")
    result.Title = "Custom Test Title"
    
    step := domain.NewTestStep("Custom Step")
    step.SetExpectedResult("Step completed")
    step.SetStatus(domain.StepStatusPassed)
    
    result.AddStep(step)
    
    if err := reporter.AddResult(result); err != nil {
        t.Fatal(err)
    }
}
```

### Test with Attachments

```go
func TestWithScreenshots(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        DisplayName: "Screenshot Test",
        Title: "Test with Screenshots",
    }, func() {
        // Create attachment
        screenshot := domain.Attachment{
            ID:       "screenshot-1",
            FileName: "homepage.png",
            MimeType: "image/png",
            Content:  []byte("fake image data"),
            Size:     16,
        }
        
        // Add attachment to result (handled automatically by qase.Test)
    })
}
```

### Test with Enhanced Error Details

```go
func TestWithErrorDetails(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        DisplayName: "Error Details Test",
        Title: "Test Error Details Capture",
        Description: "Demonstrates enhanced error information capture",
    }, func() {
        // This test will fail and capture detailed error information
        userID := "12345"
        expectedUserID := "67890"
        
        // When this assertion fails, detailed error information is captured
        qase.Equal(t, expectedUserID, userID, "User ID should match expected value")
        
        // The resulting TestResult will contain:
        // - ErrorDetails.ErrorType = "Not equal"
        // - ErrorDetails.Expected = "67890"
        // - ErrorDetails.Actual = "12345"
        // - ErrorDetails.File = "/path/to/test.go"
        // - ErrorDetails.Line = line number where assertion failed
    })
}
```

### Accessing Error Details Programmatically

```go
func TestAccessErrorDetails(t *testing.T) {
    result := domain.NewTestResult("Error Details Access Test")
    
    // Set current test result for error capture
    qase.setCurrentTestResult(result)
    defer qase.clearCurrentTestResult()
    
    // Simulate a failed assertion
    qase.True(t, false, "This will fail")
    
    // Access captured error details
    if result.Execution.ErrorDetails != nil {
        fmt.Printf("Error Type: %s\n", result.Execution.ErrorDetails.ErrorType)
        fmt.Printf("Expected: %s\n", result.Execution.ErrorDetails.Expected)
        fmt.Printf("Actual: %s\n", result.Execution.ErrorDetails.Actual)
        fmt.Printf("File: %s\n", result.Execution.ErrorDetails.File)
        fmt.Printf("Line: %d\n", result.Execution.ErrorDetails.Line)
    }
}
```

## Status Mapping

The SDK maps test statuses to Qase TestOps statuses:

| SDK Status | Qase Status |
|------------|-------------|
| `domain.StatusPassed` | "passed" |
| `domain.StatusFailed` | "failed" |
| `domain.StatusBlocked` | "blocked" |
| `domain.StatusSkipped` | "skipped" |
| `domain.StatusInProgress` | "in_progress" |
| `domain.StatusInvalid` | "invalid" |

## Best Practices

1. **Use Descriptive Titles**: Make test titles clear and descriptive
2. **Add Descriptions**: Provide detailed descriptions for complex tests
3. **Use Steps**: Break down tests into logical steps
4. **Handle Errors**: Always handle errors appropriately
5. **Use Metadata**: Add relevant metadata for better organization
6. **Use Attachments**: Add screenshots and logs when helpful
7. **Configure Fallbacks**: Always configure fallback reporters for reliability

## Integration with Qase TestOps

The SDK integrates seamlessly with Qase TestOps:

1. **Configuration**: Set up with TestOps credentials
2. **Test Execution**: Run tests with automatic reporting
3. **Result Upload**: Results are automatically uploaded to TestOps
4. **Dashboard**: View results in Qase TestOps dashboard

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
