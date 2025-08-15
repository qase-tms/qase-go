# Qase Go SDK

Qase Go SDK provides comprehensive tools for integrating Go applications with Qase TestOps. The SDK includes domain models, reporters, and testing utilities.

## Project Structure

This repository contains multiple Go modules:

- **Root Module** (`github.com/qase-tms/qase-go`): Main project module
- **Package Module** (`github.com/qase-tms/qase-go/pkg/qase-go`): Main SDK package for external use
- **API Client Modules**: Generated API clients for Qase TestOps
- **Examples**: Standalone examples demonstrating SDK usage

## Installation

To use the SDK in your project:

```bash
go get github.com/qase-tms/qase-go/pkg/qase-go
```

## Quick Start

```go
package main

import (
    "context"
    "testing"
    
    "github.com/qase-tms/qase-go/pkg/qase-go/config"
    "github.com/qase-tms/qase-go/pkg/qase-go/domain"
    "github.com/qase-tms/qase-go/pkg/qase-go/reporters"
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
    reporter, err := reporters.NewCoreReporter(cfg)
    if err != nil {
        t.Fatal(err)
    }
    
    // Start test run
    ctx := context.Background()
    if err := reporter.StartTestRun(ctx); err != nil {
        t.Fatal(err)
    }
    defer reporter.CompleteTestRun(ctx)
    
    // Create test result
    result := domain.NewTestResult("My Feature Test")
    result.Title = "My Feature Test"
    result.Fields["description"] = "Testing my feature functionality"
    
    // Create test step
    step := domain.NewTestStep("Perform action")
    step.SetExpectedResult("Action completed successfully")
    step.SetStatus(domain.StepStatusPassed)
    
    result.AddStep(*step)
    
    // Add result to reporter
    if err := reporter.AddResult(result); err != nil {
        t.Fatal(err)
    }
}
```

## Features

- **Domain Models**: Complete Go representations of Qase TestOps data structures
- **Reporters**: Flexible reporting system with TestOps and file-based reporters
- **Testing Utilities**: Helper functions for creating test reports
- **Configuration Management**: Easy configuration setup for different environments
- **Error Handling**: Comprehensive error handling and validation

## Documentation

For detailed documentation and examples, see:
- [Package Documentation](./qase-go/pkg/qase-go/README.md) - Complete SDK documentation
- [Examples](./examples/simple-examples/) - Standalone usage examples

## Development

This project uses Go workspaces for managing multiple modules:

```bash
# Initialize workspace
go work init . ./pkg/qase-go ./qase-api-client ./qase-api-v2-client

# Sync dependencies
go work sync

# Build packages
go build ./pkg/qase-go/...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
