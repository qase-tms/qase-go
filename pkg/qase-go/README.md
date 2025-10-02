# Qase Go SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/qase-tms/qase-go)](https://goreportcard.com/report/github.com/qase-tms/qase-go)
[![GoDoc](https://godoc.org/github.com/qase-tms/qase-go?status.svg)](https://godoc.org/github.com/qase-tms/qase-go)

Qase Go SDK provides comprehensive tools for integrating Go applications with Qase TestOps. The SDK includes domain models, reporters, and testing utilities for seamless test result reporting.

## Features

- **Simple Integration**: Easy setup with minimal configuration
- **Flexible Reporting**: Support for TestOps and local file reporting
- **Rich Assertions**: Comprehensive assertion library with detailed error capture
- **Test Organization**: Suite management and test metadata support
- **Configuration Management**: Environment-based configuration with fallback support
- **Status Mapping**: Transform test result statuses through configuration

## Installation

```bash
go get github.com/qase-tms/qase-go/pkg/qase-go
```

**⚠️ Important**: After installation, you must add `TestMain` to each package for the reporter to work properly. See the [Quick Start](#quick-start) section below.

## Quick Start

### Basic Usage

**Important**: Add `TestMain` to each package to initialize the reporter:

```go
package main

import (
    "testing"
    "github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

// TestMain ensures that results are sent for this package
func TestMain(m *testing.M) {
    qase.TestMainForPackage(m)
}

func TestMyFeature(t *testing.T) {
    qase.Test(t, qase.TestMetadata{
        Title:       "My Feature Test",
        Description: "Testing my feature functionality",
    }, func() {
        // Your test logic here
        qase.True(t, true, "This should pass")
    })
}
```

### Configuration

Create a `qase.config.json` file in your project root:

```json
{
  "mode": "testops",
  "fallback": "report",
  "debug": false,
  "logging": {
    "console": true,
    "file": true
  },
  "statusMapping": {
    "invalid": "failed",
    "skipped": "passed"
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
    "statusFilter": ["passed", "failed"]
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

### TestMain Setup

Add this to each package to initialize the reporter:

```go
// TestMain ensures that results are sent for this package
func TestMain(m *testing.M) {
    qase.TestMainForPackage(m)
}
```

Or use environment variables:

```bash
export QASE_MODE=testops
export QASE_TESTOPS_API_TOKEN=your-api-token
export QASE_TESTOPS_PROJECT=your-project-code
export QASE_TESTOPS_RUN_ID=123
```

### Running Tests

**Note**: Make sure you have added `TestMain` to each package before running tests.

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
qasectl testops run create --project PROJ --token <token> --title "Test Run 1"
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
qasectl testops run complete --project PROJ --token <token> --id 1
```

**GitHub Actions:**

If you're using GitHub, we provide ready-to-use actions:

- [Create test runs](https://github.com/qase-tms/gh-actions/run-create)
- [Complete test runs](https://github.com/qase-tms/gh-actions/run-complete)

#### Test Execution

```bash
# Run tests with reporting
go test ./...

# Run specific test
go test -run TestMyFeature ./...

# Run with verbose output
go test -v ./...
```

## Configuration Options

| Environment Variable | Description | Default |
|---------------------|-------------|---------|
| `QASE_MODE` | Reporting mode: `testops`, `report`, or `off` | `off` |
| `QASE_FALLBACK` | Fallback mode if main mode fails | `off` |
| `QASE_DEBUG` | Enable debug logging | `false` |
| `QASE_ENVIRONMENT` | Environment name | `local` |
| `QASE_CAPTURE_LOGS` | Capture test logs | `false` |
| `QASE_TESTOPS_API_TOKEN` | Qase API token | - |
| `QASE_TESTOPS_PROJECT` | Project code | - |
| `QASE_TESTOPS_RUN_ID` | Test run ID | - |
| `QASE_TESTOPS_API_HOST` | API host | `qase.io` |
| `QASE_TESTOPS_DEFECT` | Auto-create defects for failed tests | `false` |
| `QASE_TESTOPS_BATCH_SIZE` | Batch size for result uploads | `100` |
| `QASE_TESTOPS_STATUS_FILTER` | Comma-separated list of statuses to exclude from TestOps | - |
| `QASE_STATUS_MAPPING` | Status mapping in format "from=to,from2=to2" | - |
| `QASE_REPORT_DRIVER` | Report driver | `local` |
| `QASE_REPORT_CONNECTION_PATH` | Local report path | `./build/qase-report` |
| `QASE_REPORT_CONNECTION_FORMAT` | Report format | `json` |
| `QASE_LOGGING_CONSOLE` | Enable console logging | `true` |
| `QASE_LOGGING_FILE` | Enable file logging | `true` |

## Documentation

For detailed documentation and advanced usage examples, see:

- [Usage Guide](docs/usage.md) - Comprehensive usage documentation
- [API Reference](https://godoc.org/github.com/qase-tms/qase-go) - Go package documentation

## License

This project is licensed under the MIT License - see the [LICENSE](../../LICENSE) file for details.
