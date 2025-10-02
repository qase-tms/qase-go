# Usage Guide

This guide provides comprehensive documentation for using the Qase Go SDK in your testing projects.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Basic Usage](#basic-usage)
- [Test Metadata](#test-metadata)
- [Step Metadata](#step-metadata)
- [Assertions](#assertions)
- [Test Run Management](#test-run-management)
- [Advanced Features](#advanced-features)
- [Best Practices](#best-practices)

## Installation

### Go Modules

```bash
go get github.com/qase-tms/qase-go/pkg/qase-go
```

### Import

```go
import "github.com/qase-tms/qase-go/pkg/qase-go/qase"
```

## Configuration

The SDK supports multiple configuration methods with flexible fallback options.

### Configuration File

Create a `qase.config.json` file in your project root:

```json
{
  "mode": "testops",
  "fallback": "report",
  "debug": false,
  "environment": "staging",
  "captureLogs": true,
  "logging": {
    "console": true,
    "file": true
  },
  "testops": {
    "api": {
      "token": "your-api-token",
      "host": "qase.io"
    },
    "project": "your-project-code",
    "run": {
      "id": 123
    },
    "defect": false,
    "batch": {
      "size": 100
    },
    "statusFilter": ["passed", "skipped"]
  },
  "report": {
    "driver": "local",
    "connection": {
      "local": {
        "path": "./build/qase-report",
        "format": "json"
      }
    }
  }
}
```

**Configuration Options:**

- `mode`: "testops", "report", or "off"
- `fallback`: Fallback mode if main mode fails
- `debug`: Enable debug logging
- `environment`: Environment name
- `captureLogs`: Capture test logs
- `logging.console`: Enable console logging (default: true)
- `logging.file`: Enable file logging (default: true)
- `testops.api.host`: "qase.io" or "api.qase.io" for self-hosted
- `testops.run.id`: **Required** - specific test run ID (must be created beforehand)
- `testops.defect`: Auto-create defects for failed tests
- `testops.batch.size`: Batch size for result uploads
- `testops.statusFilter`: Array of test result statuses to exclude from TestOps reporting
- `report.connection.local.format`: "json" or "xml"

**Note**: For TestOps mode, you must create a test run first and specify its ID in the configuration. See [Test Run Management](#test-run-management) section below.

### Environment Variables

```bash
# Main Configuration
export QASE_MODE=testops
export QASE_FALLBACK=report
export QASE_DEBUG=false
export QASE_ENVIRONMENT=staging
export QASE_CAPTURE_LOGS=true

# Logging Configuration
export QASE_LOGGING_CONSOLE=true
export QASE_LOGGING_FILE=true

# TestOps Configuration
export QASE_TESTOPS_API_TOKEN=your-api-token
export QASE_TESTOPS_PROJECT=your-project-code
export QASE_TESTOPS_RUN_ID=123
export QASE_TESTOPS_API_HOST=qase.io
export QASE_TESTOPS_DEFECT=false
export QASE_TESTOPS_BATCH_SIZE=100
export QASE_TESTOPS_STATUS_FILTER=passed,skipped

# Report Configuration
export QASE_REPORT_DRIVER=local
export QASE_REPORT_CONNECTION_PATH=./build/qase-report
export QASE_REPORT_CONNECTION_FORMAT=json
```

### Programmatic Configuration

```go
import "github.com/qase-tms/qase-go/pkg/qase-go/config"

cfg := &config.Config{
    Mode:     "testops",
    Fallback: "report",
    Debug:    true,
    Logging: config.LoggingConfig{
        Console: true,
        File:    true,
    },
    TestOps: config.TestOpsConfig{
        API: config.APIConfig{
            Token: "your-api-token",
            Host:  "qase.io",
        },
        Project: "your-project-code",
        Run: config.RunConfig{
            ID: &runID, // Optional
        },
    },
    Report: config.ReportConfig{
        Driver: "local",
        Connection: config.ConnectionConfig{
            Local: config.LocalConfig{
                Path:   "./reports",
                Format: "json",
            },
        },
    },
}

// Initialize with custom config
err := qase.InitializeGlobalWithConfig(cfg)
if err != nil {
    log.Fatal(err)
}
```

## Basic Usage

### Package Setup

First, add `TestMain` to each package to initialize the reporter:

```go
// TestMain ensures that results are sent for this package
func TestMain(m *testing.M) {
    qase.TestMainForPackage(m)
}
```

### Simple Test

```go
func TestLoginFlow(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        Title:       "User Login Flow",
        Description: "Test the complete user login process",
    }, func() {
        // Test logic here
        qase.True(t, true, "Login should succeed")
    })
}
```

### Test with Steps

```go
func TestCheckoutProcess(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        Title:       "Checkout Process",
        Description: "Test the complete checkout flow",
    }, func() {
        qase.Step(t, qase.StepMetadata{
            Name:           "Add item to cart",
            ExpectedResult: "Item should be added to cart successfully",
        }, func() {
            qase.AddMessage("Adding product to shopping cart")
            qase.True(t, addItemToCart(), "Item should be added to cart")
        })
        
        qase.Step(t, qase.StepMetadata{
            Name:           "Proceed to checkout",
            ExpectedResult: "User should be redirected to checkout page",
        }, func() {
            qase.AddMessage("Navigating to checkout")
            qase.True(t, proceedToCheckout(), "Should proceed to checkout")
        })
        
        qase.Step(t, qase.StepMetadata{
            Name:           "Complete payment",
            ExpectedResult: "Payment should be processed successfully",
        }, func() {
            qase.AddMessage("Processing payment")
            qase.True(t, completePayment(), "Payment should be completed")
        })
    })
}
```

## Test Metadata

The `TestMetadata` struct provides comprehensive test information:

```go
type TestMetadata struct {
    // Core test information
    Title       string // Test title
    Description string // Test description
    Comment     string // Test comment
    
    // Suite information
    Suite []string // Test suite hierarchy
    
    // Metadata fields
    Fields map[string]string // Custom fields
    
    // Parameters
    Parameters map[string]string // Test parameters
    
    // Group parameters
    GroupParameters map[string]string // Group parameters
    
    // Test control
    Ignore bool // Skip this test
}
```

## Step Metadata

The `StepMetadata` struct provides information for test steps:

```go
type StepMetadata struct {
    Name           string // Step name
    ExpectedResult string // Expected result description
    Data           string // Step data/input
}
```

### Step Examples

```go
// Basic step
qase.Step(t, qase.StepMetadata{
    Name:           "Login with valid credentials",
    ExpectedResult: "User should be logged in successfully",
}, func() {
    qase.AddMessage("Entering username and password")
    qase.True(t, login("user", "pass"), "Login should succeed")
})

// Step with data
qase.Step(t, qase.StepMetadata{
    Name:           "Process payment",
    ExpectedResult: "Payment should be processed successfully",
    Data:           "amount: 100, currency: USD",
}, func() {
    qase.AddMessage("Processing payment of $100")
    qase.True(t, processPayment(100, "USD"), "Payment should complete")
})
```

### Examples

```go
// Basic metadata
qase.Test(t, qase.TestMetadata{
    Title: "API Authentication Test",
}, func() {
    // Test logic
})

// Rich metadata
qase.Test(t, qase.TestMetadata{
    Title:       "User Registration Flow",
    Description: "Test user registration with email verification",
    Comment:     "Registration Flow Test",
    Suite:       []string{"Authentication", "Registration"},
    Fields: map[string]string{
        "environment": "staging",
        "browser":     "chrome",
        "version":     "1.0.0",
    },
    Parameters: map[string]string{
        "test_type": "integration",
    },
    GroupParameters: map[string]string{
        "group": "auth_tests",
    },
}, func() {
    // Test logic
})

// Nested suites
qase.Test(t, qase.TestMetadata{
    Title: "Payment Processing",
    Suite: []string{"E-commerce", "Payment", "Credit Card"},
}, func() {
    // Test logic
})
```

## Assertions

The SDK provides a comprehensive assertion library with detailed error capture.

### Basic Assertions

```go
// Boolean assertions
qase.True(t, condition, "message")
qase.False(t, condition, "message")

// Equality assertions
qase.Equal(t, expected, actual, "message")
qase.NotEqual(t, expected, actual, "message")
qase.EqualValues(t, expected, actual, "message")
qase.NotEqualValues(t, expected, actual, "message")
```

### Error Assertions

```go
// Error handling
qase.Error(t, err, "message")
qase.NoError(t, err, "message")
qase.EqualError(t, err, expectedMessage, "message")
qase.ErrorIs(t, err, target, "message")
qase.ErrorAs(t, err, target, "message")
```

### String Assertions

```go
// String operations
qase.Contains(t, str, substr, "message")
qase.NotContains(t, str, substr, "message")
qase.Regexp(t, rx, str, "message")
qase.NotRegexp(t, rx, str, "message")
```

### Numeric Assertions

```go
// Numeric comparisons
qase.Greater(t, e1, e2, "message")
qase.GreaterOrEqual(t, e1, e2, "message")
qase.Less(t, e1, e2, "message")
qase.LessOrEqual(t, e1, e2, "message")
qase.Zero(t, value, "message")
qase.NotZero(t, value, "message")
```

### Collection Assertions

```go
// Collection operations
qase.Empty(t, object, "message")
qase.NotEmpty(t, object, "message")
qase.Len(t, object, length, "message")
qase.Same(t, expected, actual, "message")
qase.NotSame(t, expected, actual, "message")
qase.Subset(t, list, subset, "message")
qase.NotSubset(t, list, subset, "message")
qase.ElementsMatch(t, listA, listB, "message")
```

### Enhanced Error Details

All assertions automatically capture detailed error information:

```go
func TestErrorDetails(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        Title: "Error Details Test",
    }, func() {
        userID := "12345"
        expectedUserID := "67890"
        
        // When this assertion fails, detailed error information is captured
        qase.Equal(t, expectedUserID, userID, "User ID should match")
        
        // The resulting TestResult will contain:
        // - ErrorDetails.ErrorType = "Not equal"
        // - ErrorDetails.Expected = "67890"
        // - ErrorDetails.Actual = "12345"
        // - ErrorDetails.File = "/path/to/test.go"
        // - ErrorDetails.Line = line number
    })
}
```

## Advanced Features

### Add Messages

Use `qase.AddMessage` to add informational messages to test steps:

```go
qase.Step(t, qase.StepMetadata{
    Name:           "User Registration",
    ExpectedResult: "User should be registered successfully",
}, func() {
    qase.AddMessage("Filling registration form")
    qase.AddMessage("Submitting form data")
    qase.True(t, registerUser("john@example.com", "password123"), "Registration should succeed")
})
```

### Status Filtering

The SDK supports filtering test results by status when reporting to TestOps. This allows you to exclude certain test result statuses from being sent to Qase TestOps, which can be useful for reducing noise in your test reports or focusing on specific types of results.

#### Status Filter Configuration

Add the `statusFilter` field to your TestOps configuration:

```json
{
  "mode": "testops",
  "testops": {
    "api": {
      "token": "your-api-token",
      "host": "qase.io"
    },
    "project": "your-project-code",
    "run": {
      "id": 123
    },
    "statusFilter": ["passed", "failed"]
  }
}
```

#### Environment Variable

You can also configure status filtering using environment variables:

```bash
export QASE_TESTOPS_STATUS_FILTER=passed,failed
```

#### Supported Statuses

The following test result statuses can be filtered:

- `passed` - Test passed successfully
- `failed` - Test failed
- `blocked` - Test was blocked
- `skipped` - Test was skipped
- `in_progress` - Test is in progress
- `invalid` - Test result is invalid

#### How It Works

- **Exclusion Logic**: Statuses listed in `statusFilter` are **excluded** from TestOps reporting
- **Case Insensitive**: Status comparison is case-insensitive
- **Empty Filter**: If `statusFilter` is empty or not specified, all results are sent to TestOps
- **Debug Logging**: When debug mode is enabled, filtered results are logged

#### Configuration Examples

**Exclude only passed tests:**

```json
{
  "testops": {
    "statusFilter": ["passed"]
  }
}
```

This will send `failed`, `blocked`, `skipped`, `in_progress`, and `invalid` results to TestOps, but exclude `passed` results.

**Exclude multiple statuses:**

```json
{
  "testops": {
    "statusFilter": ["passed", "skipped"]
  }
}
```

This will send only `failed`, `blocked`, `in_progress`, and `invalid` results to TestOps.

**Exclude all statuses:**

```json
{
  "testops": {
    "statusFilter": ["passed", "failed", "blocked", "skipped", "in_progress", "invalid"]
  }
}
```

This will prevent any results from being sent to TestOps.

#### Use Cases

- **Focus on Failures**: Exclude `passed` and `skipped` results to focus only on test failures
- **Reduce Noise**: Exclude `skipped` tests from reports to reduce clutter
- **Debug Mode**: Exclude `passed` tests during debugging to focus on problematic areas

### Multi-package Testing

For multi-package projects, add `TestMain` to each package to ensure proper initialization and result reporting:

```go
// TestMain ensures that results are sent for this package
func TestMain(m *testing.M) {
    qase.TestMainForPackage(m)
}
```

**Important**: Add this function **once per package** to initialize the reporter and ensure test results are properly sent to Qase TestOps or saved to local files.

### Test Run Management

**Important**: The reporter does not create or complete test runs automatically. You must specify an existing test run ID in your configuration.

#### Creating Test Runs

You can create test runs using:

**cURL:**

```bash
curl --request POST \
     --url https://api.qase.io/v1/run/DEMO \
     --header 'Token: YOUR_TOKEN' \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "title": "My test run"
}
'
```

**qasectl CLI:**

```bash
qasectl run create --project DEMO --title "My test run"
```

#### Completing Test Runs

You can complete test runs using:

**cURL:**

```bash
curl --request POST \
     --url https://api.qase.io/v1/run/DEMO/1/complete \
     --header 'Token: YOUR_TOKEN' \
     --header 'accept: application/json'
```

**qasectl CLI:**

```bash
qasectl run complete --project DEMO --id 1
```

**GitHub Actions:**

If you're using GitHub, we provide ready-to-use actions:

- [Create test runs](https://github.com/qase-tms/gh-actions/run-create)
- [Complete test runs](https://github.com/qase-tms/gh-actions/run-complete)

### Attachments

Add screenshots, logs, and other files to test results:

```go
func TestWithAttachments(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        Title:       "Test with Attachments",
        Description: "Demonstrates adding attachments to test results",
    }, func() {
        qase.AddMessage("Testing with attachments")
        
        // Add file attachment
        qase.AddAttachments("../go.mod")
        
        // Add content attachment
        qase.AttachContent("log.txt", "My logs", "plain/text")
        
        qase.True(t, true)
    })
}
```

## Best Practices

### 1. Always Add TestMain

Add `TestMain` to each package to ensure proper initialization:

```go
// TestMain ensures that results are sent for this package
func TestMain(m *testing.M) {
    qase.TestMainForPackage(m)
}
```

### 2. Use Descriptive Test Names

```go
// Good
func TestUserLoginWithValidCredentials(t *testing.T) {
    // Test logic
}

// Avoid
func TestLogin(t *testing.T) {
    // Test logic
}
```

### 3. Add Comprehensive Metadata

```go
qase.Test(t, qase.TestMetadata{
    Title:       "User Registration with Email Verification",
    Description: "Test the complete user registration flow including email verification step",
    Comment:     "Critical registration test",
    Suite:       []string{"Authentication", "Registration"},
    Fields: map[string]string{
        "environment": "staging",
        "browser":     "chrome",
        "version":     "1.0.0",
    },
    Parameters: map[string]string{
        "test_type": "critical",
    },
}, func() {
    // Test logic
})
```

### 4. Use Steps for Complex Tests

```go
qase.Test(t, qase.TestMetadata{
    Title: "E-commerce Checkout Flow",
}, func() {
    qase.Step(t, qase.StepMetadata{
        Name:           "Add item to cart",
        ExpectedResult: "Item should be added to cart successfully",
    }, func() {
        qase.AddMessage("Adding product to shopping cart")
        qase.True(t, addItemToCart(), "Item should be added")
    })
    
    qase.Step(t, qase.StepMetadata{
        Name:           "Proceed to checkout",
        ExpectedResult: "User should be redirected to checkout page",
    }, func() {
        qase.AddMessage("Navigating to checkout")
        qase.True(t, proceedToCheckout(), "Should proceed to checkout")
    })
    
    qase.Step(t, qase.StepMetadata{
        Name:           "Complete payment",
        ExpectedResult: "Payment should be processed successfully",
    }, func() {
        qase.AddMessage("Processing payment")
        qase.True(t, completePayment(), "Payment should complete")
    })
})
```

### 5. Handle Errors Appropriately

```go
func TestWithErrorHandling(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        Title: "API Error Handling",
    }, func() {
        // Test error scenarios
        err := callAPI()
        qase.Error(t, err, "API should return error for invalid input")
        qase.EqualError(t, err, "invalid input", "Error message should match")
    })
}
```

### 6. Use Attachments for Debugging

```go
func TestWithScreenshots(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        Title: "UI Test with Screenshots",
    }, func() {
        // Take screenshot on failure
        if !performUIAction() {
            qase.AddMessage("UI action failed, taking screenshot")
            qase.AddAttachments("screenshot.png")
            qase.AttachContent("error.log", "UI action failed at step 3", "text/plain")
        }
    })
}
```

### 7. Configure Fallbacks

```go
cfg := &config.Config{
    Mode:     "testops",    // Primary mode
    Fallback: "report",     // Fallback to local reporting
    // ... other config
}
```

### 8. Configure Logging

The SDK provides flexible logging configuration to control where logs are output:

**File-based configuration:**

```json
{
  "debug": true,
  "logging": {
    "console": true,
    "file": false
  }
}
```

**Environment variables:**

```bash
# Disable console logging
export QASE_LOGGING_CONSOLE=false

# Enable debug mode
export QASE_DEBUG=true
```

**Logging scenarios:**

- **Default**: Both console and file enabled
- **Debug mode**: All logs including debug information
- **File only**: Logs written to file, no console output
- **Console only**: Logs written to console, no file output
- **Disabled**: No logging output
