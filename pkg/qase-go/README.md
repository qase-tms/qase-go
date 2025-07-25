# Qase Go SDK

Qase Go SDK provides comprehensive tools for integrating Go applications with Qase TestOps. The SDK includes domain models, reporters, and a testing framework inspired by Allure.

## Features

- **Domain Models**: Complete Go representations of Qase TestOps data structures
- **Reporters**: Flexible reporting system with TestOps and file-based reporters
- **Testing Framework**: Fluent interface for creating test reports (inspired by Allure)
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
    "github.com/qase-tms/qase-go/pkg/qase-go/framework"
    "github.com/qase-tms/qase-go/pkg/qase-go/reporters"
)

func TestMyFeature(t *testing.T) {
    // Create configuration
    cfg := &config.Config{
        Mode: "report",
        Report: config.ReportConfig{
            Path: "./reports",
        },
    }
    
    // Create reporter
    coreReporter, err := reporters.NewCoreReporter(cfg)
    if err != nil {
        t.Fatal(err)
    }
    
    reporter := framework.NewReporter(coreReporter)
    
    // Start test run
    ctx := context.Background()
    if err := reporter.StartTestRun(ctx); err != nil {
        t.Fatal(err)
    }
    defer reporter.CompleteTestRun(ctx)
    
    // Run test with framework
    framework.TestWithReporter(t, "My Feature Test", reporter, func(t framework.Provider) {
        t.Title("My Feature Test")
        t.Description("Testing my feature functionality")
        
        // Create step with builder pattern
        step := framework.NewStep("Perform action").
            WithExpectedResult("Action completed successfully").
            WithData("input: test data").
            Passed().
            Build()
        t.Step(step)
    })
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

### Framework

The framework package provides a fluent interface for creating test reports, inspired by Allure:

- **Provider Interface**: Test execution management
- **StepBuilder**: Fluent interface for creating test steps
- **TestRunner**: Automatic test execution and reporting
- **Suite Support**: Test suite organization

```go
import "github.com/qase-tms/qase-go/pkg/qase-go/framework"

// Simple test
framework.Test(t, "My Test", func(t framework.Provider) {
    t.Title("My Test")
    
    step := framework.NewStep("Action").
        WithExpectedResult("Result").
        Passed().
        Build()
    t.Step(step)
})

// Test with timing
step := framework.NewStep("Timed Action").
    Begin().
    WithExpectedResult("Completed")
// ... do work ...
step = step.Finish().Passed()
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
            BaseURL: "https://api.qase.io",
        },
        Project: "your-project-code",
    },
    
    // Report configuration
    Report: config.ReportConfig{
        Path: "./reports",
        Format: "json",
    },
}
```

## Examples

### Basic Test with Framework

```go
func TestLoginFlow(t *testing.T) {
    framework.Test(t, "Login Flow Test", func(t framework.Provider) {
        t.Title("User Login Flow")
        t.Description("Test the complete user login process")
        
        // Set metadata
        t.SetField("environment", "staging")
        t.SetField("browser", "chrome")
        
        // Navigate to login page
        step1 := framework.NewStep("Navigate to login page").
            WithExpectedResult("Login page is displayed").
            Passed().
            Build()
        t.Step(step1)
        
        // Enter credentials
        step2 := framework.NewStep("Enter credentials").
            WithExpectedResult("Credentials are entered").
            WithData("username: admin, password: secret").
            Passed().
            Build()
        t.Step(step2)
        
        // Submit form
        step3 := framework.NewStep("Submit login form").
            WithExpectedResult("User is logged in").
            Passed().
            Build()
        t.Step(step3)
    })
}
```

### Test with Nested Steps

```go
func TestRegistrationFlow(t *testing.T) {
    framework.Test(t, "Registration Flow", func(t framework.Provider) {
        t.Title("User Registration")
        
        // Parent step
        parentStep := framework.NewStep("Complete registration").
            WithExpectedResult("User is registered").
            Begin()
        
        // Child steps
        child1 := framework.NewStep("Fill form").
            WithExpectedResult("Form is filled").
            Passed().
            Build()
        
        child2 := framework.NewStep("Verify email").
            WithExpectedResult("Email verified").
            Passed().
            Build()
        
        parentStep = parentStep.WithChild(child1).WithChild(child2)
        parentStep = parentStep.Finish().Passed()
        t.Step(parentStep.Build())
    })
}
```

### Test with Attachments

```go
func TestWithScreenshots(t *testing.T) {
    framework.Test(t, "Screenshot Test", func(t framework.Provider) {
        t.Title("Test with Screenshots")
        
        // Create attachment
        screenshot := domain.Attachment{
            ID:       "screenshot-1",
            FileName: "homepage.png",
            MimeType: "image/png",
            Content:  []byte("fake image data"),
            Size:     16,
        }
        
        t.AddAttachment(screenshot)
        
        // Step with attachment
        step := framework.NewStep("Take screenshot").
            WithExpectedResult("Screenshot captured").
            WithAttachment(screenshot).
            Passed().
            Build()
        t.Step(step)
    })
}
```

### Test Suite

```go
type MyTestSuite struct {
    framework.BaseSuite
}

func (s *MyTestSuite) BeforeEach(t framework.Provider) {
    t.SetField("suite", "MyTestSuite")
}

func (s *MyTestSuite) TestFeature(t framework.Provider) {
    t.Title("My Feature")
    
    step := framework.NewStep("Test feature").
        WithExpectedResult("Feature works").
        Passed().
        Build()
    t.Step(step)
}

func TestMySuite(t *testing.T) {
    suite := &MyTestSuite{}
    framework.TestWithSuite(t, "Suite Test", suite, func(t framework.Provider) {
        suite.TestFeature(t)
    })
}
```

## Status Mapping

The framework maps test statuses to Qase TestOps statuses:

| Framework Status | Qase Status |
|------------------|-------------|
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
7. **Use Suites**: Organize related tests into suites
8. **Configure Fallbacks**: Always configure fallback reporters for reliability

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
