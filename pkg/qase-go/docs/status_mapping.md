# Status Mapping

The Qase Go SDK supports status mapping to transform test result statuses before reporting. This feature allows you to standardize statuses across different testing frameworks or adapt to specific reporting requirements.

## Overview

Status mapping is applied **centrally** in the main reporter class before any status filtering. This ensures consistent behavior across all reporters regardless of the testing framework used.

## Available Statuses

The following statuses are available for mapping:

- `passed` - Test passed successfully
- `failed` - Test failed due to assertion error  
- `skipped` - Test was skipped
- `blocked` - Test was blocked
- `invalid` - Test failed due to non-assertion error (network issues, syntax errors)

## Configuration

### JSON Configuration

Add `statusMapping` to your `qase.config.json`:

```json
{
  "mode": "testops",
  "fallback": "report",
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
    }
  }
}
```

### Environment Variable

Use the `QASE_STATUS_MAPPING` environment variable:

```bash
export QASE_STATUS_MAPPING="invalid=failed,skipped=passed"
```

The format is `"from=to,from2=to2"` with comma-separated pairs.

## Usage Examples

### Example 1: Map Invalid Tests to Failed

```json
{
  "statusMapping": {
    "invalid": "failed"
  }
}
```

This will transform all tests with `invalid` status to `failed` status before reporting.

### Example 2: Map Skipped Tests to Passed

```json
{
  "statusMapping": {
    "skipped": "passed"
  }
}
```

This will transform all tests with `skipped` status to `passed` status before reporting.

### Example 3: Multiple Mappings

```json
{
  "statusMapping": {
    "invalid": "failed",
    "skipped": "passed",
    "blocked": "failed"
  }
}
```

This will apply multiple status transformations.

### Example 4: Environment Variable with Multiple Mappings

```bash
export QASE_STATUS_MAPPING="invalid=failed,skipped=passed,blocked=failed"
```

## Programmatic Usage

You can also create and use status mappings programmatically:

```go
package main

import (
    "github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func main() {
    // Create mapping from map
    mapping, err := domain.NewStatusMapping(map[string]string{
        "invalid": "failed",
        "skipped": "passed",
    })
    if err != nil {
        panic(err)
    }

    // Apply mapping to a status
    originalStatus := domain.StatusInvalid
    mappedStatus := mapping.ApplyMapping(originalStatus)
    // mappedStatus will be domain.StatusFailed

    // Apply mapping to a test result
    result := &domain.TestResult{
        Title: "Test Case",
        Execution: domain.TestExecution{
            Status: domain.StatusInvalid,
        },
    }
    
    changed := mapping.ApplyMappingToResult(result)
    // changed will be true, result.Execution.Status will be domain.StatusFailed
}
```

## Validation

The SDK validates status mappings to ensure:

- Source and target statuses are valid
- No mapping to the same status
- Proper error messages for invalid configurations

### Validation Examples

```go
// Valid mapping
validMapping := map[string]string{
    "invalid": "failed",
    "skipped": "passed",
}
err := domain.ValidateMapping(validMapping)
// err will be nil

// Invalid mapping - unknown source status
invalidMapping := map[string]string{
    "unknown": "failed",
}
err := domain.ValidateMapping(invalidMapping)
// err will contain validation error
```

## Important Notes

- Status mapping is applied **before** status filtering
- Mapping is applied **centrally** in the main reporter class
- Works for **all reporters** regardless of testing framework
- Invalid mappings are ignored with a warning
- Status mapping is logged when debug mode is enabled
- Case insensitive - `"INVALID"` and `"invalid"` are treated the same
- Whitespace is trimmed automatically

## Error Handling

If an invalid status mapping is configured:

1. **JSON Configuration**: The configuration validation will fail with a descriptive error message
2. **Environment Variable**: Invalid mappings are ignored with a warning
3. **Programmatic Usage**: Functions return errors for invalid configurations

## Debugging

Enable debug mode to see status mapping in action:

```json
{
  "debug": true,
  "statusMapping": {
    "invalid": "failed"
  }
}
```

This will log status changes like:

```
Status mapped for test 'My Test': invalid -> failed
```

## Best Practices

1. **Use consistent mappings** across your test suite
2. **Document your mappings** for team members
3. **Test your mappings** with sample data
4. **Use environment variables** for different environments
5. **Validate configurations** before deployment

## Migration Guide

If you're upgrading from a version without status mapping:

1. No changes required - existing configurations will continue to work
2. Add `statusMapping` field to your configuration as needed
3. Test your mappings with a small subset of tests first
4. Monitor logs to ensure mappings are applied correctly
