# Qase Go SDK

Go SDK for sending test results to Qase TestOps.

## Features

- ✅ Support for both Qase API v1 and v2
- ✅ Domain models compatible with JavaScript SDK
- ✅ Automatic test run creation and management
- ✅ Batch result sending for better performance
- ✅ Environment variable configuration
- ✅ Comprehensive test result metadata support
- ✅ Attachment handling
- ✅ Test steps support
- ✅ Debug logging

## Installation

```bash
go get github.com/qase-tms/qase-go/pkg/qase-go
```

## Quick Start

### Basic Usage

The TestOps reporter has a simple 3-method API:

1. `StartTestRun()` - Creates a new test run (title/description from config)
2. `AddResult()` - Adds test results to the collection  
3. `CompleteTestRun()` - Sends all results and completes the run

```go
package main

import (
    "context"
    "log"

    "github.com/qase-tms/qase-go/pkg/qase-go/clients"
    "github.com/qase-tms/qase-go/pkg/qase-go/config"
    "github.com/qase-tms/qase-go/pkg/qase-go/domain"
    "github.com/qase-tms/qase-go/pkg/qase-go/reporters"
)

func main() {
    // Create configuration
    cfg := config.NewConfig()
    cfg.TestOps.API.Token = "your_api_token"
    cfg.TestOps.Project = "PROJECT"
    cfg.TestOps.Run.Title = "My Test Run"
    cfg.TestOps.Run.Description = "Test run description"
    cfg.TestOps.Batch.Size = 25

    // Create unified client (uses v1 for runs, v2 for results)
    client, err := clients.NewUnifiedClient(cfg)
    if err != nil {
        log.Fatal(err)
    }

    // Create reporter with unified client
    reporter := reporters.NewTestOpsReporter(client)

    ctx := context.Background()

    // 1. Start test run (title and description from config)
    err = reporter.StartTestRun(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // 2. Add test results
    result := reporters.CreateSimpleResult("My test", domain.StatusPassed)
    result.SetTestopsIDSingle(123) // Link to existing test case
    err = reporter.AddResult(result)
    if err != nil {
        log.Fatal(err)
    }

    // 3. Complete test run (sends all results and completes the run)
    err = reporter.CompleteTestRun(ctx)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Test run completed successfully!")
}
```

## Configuration

### Unified Client Configuration

The reporter uses a unified client that automatically chooses the best API for each operation:
- **API v1** for test run management (create/complete)
- **API v2** for result uploading (faster and more reliable)

The `TestOpsReporter` accepts any client that implements the `TestOpsClient` interface (defined in the reporters package following Go interface best practices).

```go
import (
    "github.com/qase-tms/qase-go/pkg/qase-go/clients"
    "github.com/qase-tms/qase-go/pkg/qase-go/config"
    "github.com/qase-tms/qase-go/pkg/qase-go/reporters"
)

// Create configuration
cfg := config.NewConfig()
cfg.TestOps.API.Token = "your_api_token"
cfg.TestOps.Project = "PROJECT"
cfg.TestOps.API.Host = "qase.io"                    // Optional: default is qase.io
cfg.TestOps.Run.Title = "My Test Run"               // Run title
cfg.TestOps.Run.Description = "Test description"    // Run description
cfg.TestOps.Batch.Size = 25                         // Batch size for results
cfg.Debug = true                                     // Optional: enable debug logging

// Create unified client
client, err := clients.NewUnifiedClient(cfg)
if err != nil {
    log.Fatal(err)
}

// Create reporter
reporter := reporters.NewTestOpsReporter(client)
```

### Environment Variables

You can also load configuration from environment variables:

```go
cfg := config.NewConfig()
cfg.LoadFromEnvironment()
// Will load QASE_TESTOPS_API_TOKEN, QASE_TESTOPS_PROJECT, etc.
```

| Variable | Description | Default |
|----------|-------------|---------|
| `QASE_TESTOPS_API_TOKEN` | Qase API token (required) | - |
| `QASE_TESTOPS_PROJECT` | Project code (required) | - |
| `QASE_TESTOPS_API_HOST` | Custom API host | qase.io |
| `QASE_TESTOPS_RUN_TITLE` | Title for new test run | "Automated Test Run" |
| `QASE_TESTOPS_RUN_DESCRIPTION` | Description for new test run | - |
| `QASE_TESTOPS_BATCH_SIZE` | Batch size for result uploading | 50 |
| `QASE_DEBUG` | Enable debug logging | false |

## Creating Test Results

### Simple Test Result

```go
result := reporters.CreateSimpleResult("Test name", domain.StatusPassed)
result.SetTestopsIDSingle(123) // Link to test case ID 123
result.SetMessage("Test completed successfully")
```

### Test Result with Steps

```go
steps := []domain.TestStep{
    reporters.CreateStep("Open browser", domain.StepStatusPassed),
    reporters.CreateStep("Navigate to page", domain.StepStatusPassed),
    reporters.CreateStep("Verify content", domain.StepStatusFailed),
}

result := reporters.CreateResultWithSteps("E2E Test", domain.StatusFailed, steps)
```

### Advanced Test Result

```go
result := domain.NewTestResult("Advanced test")
result.Execution.Status = domain.StatusPassed
result.SetTestopsIDSingle(456)

// Set timing
now := time.Now().Unix()
start := now - 30 // 30 seconds ago
result.Execution.StartTime = &start
result.Execution.EndTime = &now
duration := int64(30000) // 30 seconds in milliseconds
result.Execution.Duration = &duration

// Add metadata
result.SetField("severity", "high")
result.SetField("priority", "critical")
result.SetParam("browser", "chrome")
result.SetParam("os", "linux")

// Add custom message
result.SetMessage("Test passed with warnings")

// Add to suite
result.AddSuiteData("Smoke Tests", nil)
```

## Test Statuses

Available test result statuses:

- `domain.StatusPassed` - Test passed
- `domain.StatusFailed` - Test failed
- `domain.StatusBlocked` - Test blocked
- `domain.StatusSkipped` - Test skipped
- `domain.StatusInProgress` - Test in progress
- `domain.StatusInvalid` - Test invalid

Available step statuses:

- `domain.StepStatusPassed` - Step passed
- `domain.StepStatusFailed` - Step failed
- `domain.StepStatusBlocked` - Step blocked
- `domain.StepStatusSkipped` - Step skipped

## Working with Steps

### Basic Steps

```go
step := reporters.CreateStep("Click login button", domain.StepStatusPassed)
result.AddStep(step)
```

### Advanced Steps

```go
step := domain.NewTestStep("Enter username")
step.Execution.Status = domain.StepStatusPassed

// Set expected result and data
expectedResult := "Username field should accept input"
step.Data.ExpectedResult = &expectedResult

data := "username: testuser@example.com"
step.Data.Data = &data

// Set timing
now := time.Now().Unix()
step.Execution.StartTime = &now
step.Execution.EndTime = &now

result.AddStep(*step)
```

### Nested Steps

```go
parentStep := domain.NewTestStep("Login process")
parentStep.Execution.Status = domain.StepStatusPassed

childStep1 := reporters.CreateStep("Enter username", domain.StepStatusPassed)
childStep2 := reporters.CreateStep("Enter password", domain.StepStatusPassed)
childStep3 := reporters.CreateStep("Click submit", domain.StepStatusPassed)

parentStep.Steps = []domain.TestStep{childStep1, childStep2, childStep3}
result.AddStep(*parentStep)
```

## API Clients

The SDK supports both Qase API v1 and v2:

### Using API v2 (Recommended)

```go
import "github.com/qase-tms/qase-go/pkg/qase-go/config"

// Option 1: Using config.Config
cfg := config.NewConfig()
cfg.TestOps.API.Token = "your_token"
cfg.Debug = true

client, err := clients.NewClientFromConfig(cfg, true) // true = use API v2
if err != nil {
    log.Fatal(err)
}

// Option 2: Using ClientConfig directly
clientConfig := clients.ClientConfig{
    APIToken: "your_token",
    UseAPIv2: true,
    Debug:    true,
}

client, err := clients.NewClient(clientConfig)
if err != nil {
    log.Fatal(err)
}

// Send single result
err = client.SendResult(ctx, "PROJECT", 123, result)

// Send multiple results
err = client.SendResults(ctx, "PROJECT", 123, []*domain.TestResult{result1, result2})
```

### Using API v1 (Limited Support)

⚠️ **Note**: API v1 client only supports test run management. For sending results, please use API v2.

```go
cfg := config.NewConfig()
cfg.TestOps.API.Token = "your_token"
cfg.Debug = true

client, err := clients.NewClientFromConfig(cfg, false) // false = use API v1

// ❌ These methods will return an error:
// client.SendResult(ctx, "PROJECT", 123, result)        // Not supported
// client.SendResults(ctx, "PROJECT", 123, results)      // Not supported

// ✅ These methods work:
// client.CreateRun(ctx, "PROJECT", "Title", "Desc")     // Supported
// client.CompleteRun(ctx, "PROJECT", runID)             // Supported
```

### Managing Test Runs

```go
// Create a new test run
runInfo, err := client.CreateRun(ctx, "PROJECT", "Test Run Title", "Description")
if err != nil {
    log.Fatal(err)
}

log.Printf("Created run: %d, URL: %s", runInfo.ID, runInfo.URL)

// Complete the test run
err = client.CompleteRun(ctx, "PROJECT", runInfo.ID)
```

## Error Handling

The SDK provides detailed error information:

```go
if err := reporter.SendResults(ctx); err != nil {
    log.Printf("Failed to send results: %v", err)
    // Handle error appropriately
}
```

Common errors:
- Invalid API token
- Project not found
- Test run not found
- Network connectivity issues
- Invalid test result data

## Performance Considerations

### Batch Processing

The reporter automatically batches results for better performance:

```go
config := reporters.TestOpsConfig{
    // ... other config
    BatchSize: 25, // Send results in batches of 25
}
```

### Periodic Sending

Enable periodic sending to reduce memory usage for long-running tests:

```go
config := reporters.TestOpsConfig{
    // ... other config
    SendInterval: 30 * time.Second, // Send results every 30 seconds
}
```

## Integration Examples

### Testing Framework Integration

```go
func TestWithQase(t *testing.T) {
    reporter, _ := reporters.NewTestOpsReporterFromEnv()
    ctx := context.Background()
    reporter.Start(ctx)
    defer reporter.Finish(ctx)

    t.Run("test case 1", func(t *testing.T) {
        // Your test logic here
        passed := true // result of your test
        
        var status domain.TestResultStatus
        if passed {
            status = domain.StatusPassed
        } else {
            status = domain.StatusFailed
        }
        
        result := reporters.CreateSimpleResult("test case 1", status)
        result.SetTestopsIDSingle(123)
        reporter.AddResult(result)
    })
}
```

### CI/CD Integration

```bash
#!/bin/bash

# Set environment variables
export QASE_TESTOPS_API_TOKEN="$QASE_API_TOKEN"
export QASE_TESTOPS_PROJECT="PROJECT"
export QASE_ENVIRONMENT="$CI_ENVIRONMENT_NAME"
export QASE_DEBUG="true"

# Run tests with Qase reporting
go run your_test_runner.go
```

## Troubleshooting

### Enable Debug Logging

```bash
export QASE_DEBUG="true"
```

Or programmatically:

```go
config := reporters.TestOpsConfig{
    Debug: true,
    // ... other config
}
```

### Common Issues

1. **API Token Issues**
   - Ensure your API token is valid and has the necessary permissions
   - Check that the token is correctly set in environment variables

2. **Project Not Found**
   - Verify the project code is correct
   - Ensure your API token has access to the project

3. **Network Issues**
   - Check internet connectivity
   - Verify firewall settings allow HTTPS traffic to api.qase.io

4. **Invalid Test Case IDs**
   - Ensure test case IDs exist in your Qase project
   - Test case IDs should be integers

## Examples

See the [examples](examples/) directory for more usage examples:

- [Simple Usage](examples/simple_usage.go) - Basic reporter usage
- More examples coming soon...

## API Reference

### Domain Models

- [`domain.TestResult`](domain/test_result.go) - Main test result structure
- [`domain.TestStep`](domain/test_step.go) - Test step structure
- [`domain.Attachment`](domain/attachment.go) - File attachment structure
- [`domain.TestResultStatus`](domain/status.go) - Test result status enum

### Reporters

- [`reporters.TestOpsReporter`](reporters/testops.go) - Main reporter interface
- [`reporters.TestOpsConfig`](reporters/testops.go) - Reporter configuration

### Clients

- [`clients.Client`](clients/client.go) - API client interface
- [`clients.V1Client`](clients/v1_client.go) - API v1 implementation
- [`clients.V2Client`](clients/v2_client.go) - API v2 implementation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.