# Multi-Package Test Structure with Global Initialization

This example demonstrates how to use qase-go with a multi-package test structure where tests are organized in different directories but share a single configuration file.

## Project Structure

```text
multi_package_demo/
├── qase.config.json          # Configuration in root
├── main_test.go              # Main package with TestMain
├── suite01/
│   └── second_test.go        # Tests in subdirectory
└── suite02/
    └── third_test.go         # Tests in another subdirectory
```

## How It Works

### 1. Global Initialization

The main package (`main_test.go`) contains a `TestMain` function that initializes qase globally:

```go
func TestMain(m *testing.M) {
    // Initialize qase globally - searches for config in parent directories
    if err := qase.InitializeGlobal(); err != nil {
        panic("Failed to initialize qase globally: " + err.Error())
    }

    // Run all tests
    _ = m.Run()

    // Complete test run automatically
    qase.TestMain(m)
}
```

### 2. Configuration Loading

The `qase.InitializeGlobal()` function:
- Searches for `qase.config.json` in the current directory and up to 5 parent directories
- Initializes the reporter once for all packages
- Sets up logging and other global state

### 3. Subpackage Tests

Tests in subpackages (`suite01`, `suite02`) don't need `TestMain` - they automatically use the globally initialized reporter:

```go
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

## Running Tests

Run all tests from the root directory:

```bash
go test -v ./...
```

Or run specific packages:

```bash
go test -v ./suite01
go test -v ./suite02
```

## Benefits

- ✅ **Single Configuration**: One `qase.config.json` file for all packages
- ✅ **Single Test Run**: All tests are reported as one test run in qase
- ✅ **No Duplication**: No need to duplicate `TestMain` in each package
- ✅ **Automatic Completion**: Test run is completed automatically after all tests finish
- ✅ **Parent Directory Search**: Configuration is found even when tests run from subdirectories

## Configuration

The `qase.config.json` file contains settings for:
- Test mode (report, testops, off)
- Logging configuration
- TestOps API settings
- Report output settings

## Key Functions

- `qase.InitializeGlobal()` - Initialize qase globally with parent directory search
- `qase.InitializeGlobalWithConfig(cfg)` - Initialize with custom configuration
- `qase.IsGlobalInitialized()` - Check if global initialization was done
- `qase.ResetGlobal()` - Reset global state (useful for testing)

## Example Output

When running tests, you'll see:

1. Global initialization in the main package
2. Configuration loaded from the root directory
3. All tests from all packages using the same reporter
4. Automatic test run completion
