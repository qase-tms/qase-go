# Global Initialization for Multi-Package Test Structure

This example demonstrates how to use qase-go with a multi-package test structure where tests are organized in different directories but share a single configuration file.

## Problem

When you have a project structure like this:

```
.
├── go.mod
├── go.sum
├── qase.config.json          # Configuration in root
├── qase_simple_test.go       # Tests in root
└── suite01/
    └── qase_second_test.go   # Tests in subdirectory
```

The tests in `suite01/` won't find the configuration file because qase searches for it relative to the current working directory.

## Solution

Use global initialization to initialize qase once for all packages:

### 1. Root Package TestMain

In your root package (where `qase.config.json` is located), add a `TestMain` function:

```go
package main

import (
    "testing"
    "github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

func TestMain(m *testing.M) {
    // Initialize qase globally - this will search for config in parent directories
    if err := qase.InitializeGlobal(); err != nil {
        panic("Failed to initialize qase globally: " + err.Error())
    }

    // Run all tests
    exitCode := m.Run()

    // Complete test run automatically
    qase.TestMain(m)
}
```

### 2. Subpackage Tests

In your subpackages, you don't need `TestMain` - just write your tests normally:

```go
package suite01

import (
    "testing"
    "github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

func TestMyTest(t *testing.T) {
    qase.Test(t,
        qase.TestMetadata{
            DisplayName: "My Test",
            Description: "A test in a subpackage",
        },
        func() {
            qase.True(t, true)
        })
}
```

## How It Works

1. **Global Initialization**: `qase.InitializeGlobal()` searches for configuration in parent directories (up to 5 levels up)
2. **Single Reporter**: All packages share the same reporter instance
3. **Automatic Completion**: The test run is completed automatically after all tests finish

## Alternative: Custom Configuration

If you need to provide a custom configuration instead of loading from file:

```go
func TestMain(m *testing.M) {
    cfg := &config.Config{
        Mode: "testops",
        TestOps: config.TestOpsConfig{
            API: config.APIConfig{
                Token: "your-token",
                Host:  "qase.io",
            },
            Project: "your-project",
        },
        // ... other config
    }

    if err := qase.InitializeGlobalWithConfig(cfg); err != nil {
        panic("Failed to initialize qase: " + err.Error())
    }

    exitCode := m.Run()
    qase.TestMain(m)
}
```

## Running Tests

Run all tests from the root directory:

```bash
go test ./...
```

Or run specific packages:

```bash
go test ./suite01
```

## Benefits

- ✅ Single configuration file for all packages
- ✅ Single test run in qase
- ✅ All tests reported together
- ✅ No need to duplicate TestMain in each package
- ✅ Automatic test run completion
