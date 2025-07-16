# Qase Go Configuration Package

This package provides a flexible and robust configuration system for Qase Go reporter, following the same patterns and environment variables as the qase-java-commons library.

## Features

- **Multiple Configuration Sources**: Support for JSON files, environment variables, and programmatic configuration
- **Priority-based Loading**: Environment variables override file configuration, which overrides defaults
- **Builder Pattern**: Fluent API for creating configurations programmatically
- **Validation**: Comprehensive configuration validation
- **Type Safety**: Strongly typed configuration structures
- **Default Values**: Sensible defaults for all configuration options

## Installation

```bash
go get github.com/qase-tms/qase-go/config
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/qase-tms/qase-go/config"
)

func main() {
    // Load configuration with defaults, file, and environment variables
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Mode: %s\n", cfg.Mode)
    fmt.Printf("Project: %s\n", cfg.TestOps.Project)
}
```

### Using Builder Pattern

```go
package main

import (
    "log"
    
    "github.com/qase-tms/qase-go/config"
)

func main() {
    cfg, err := config.NewConfigBuilder().
        TestOpsMode().
        WithAPIToken("your-api-token").
        WithProject("DEMO").
        WithDebug(true).
        AddRunTag("smoke").
        AddRunTag("regression").
        AddRunConfiguration("browser", "chrome").
        AddRunConfiguration("environment", "staging").
        WithBatchSize(50).
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    // Use configuration...
}
```

### Loading from File

```go
package main

import (
    "log"
    
    "github.com/qase-tms/qase-go/config"
)

func main() {
    // Load from specific file
    cfg, err := config.LoadFrom("./configs/qase.config.json")
    if err != nil {
        log.Fatal(err)
    }
    
    // Or use builder with file loading
    cfg, err = config.NewConfigBuilder().
        LoadFromFileIfExists("qase.config.json").
        LoadFromEnvironment().
        Build()
    
    if err != nil {
        log.Fatal(err)
    }
    
    // Use configuration...
}
```

## Configuration Structure

The configuration follows this JSON structure:

```json
{
  "mode": "testops",
  "fallback": "report",
  "debug": false,
  "environment": "local",
  "captureLogs": false,
  "report": {
    "driver": "local",
    "connection": {
      "local": {
        "path": "./build/qase-report",
        "format": "json"
      }
    }
  },
  "testops": {
    "api": {
      "token": "<token>",
      "host": "qase.io"
    },
    "run": {
      "title": "Regress run",
      "description": "Regress run description",
      "complete": true,
      "tags": ["tag1", "tag2"],
      "configurations": {
        "values": [
          {
            "name": "browser",
            "value": "chrome"
          },
          {
            "name": "environment",
            "value": "staging"
          }
        ],
        "createIfNotExists": true
      }
    },
    "defect": false,
    "project": "<project_code>",
    "batch": {
      "size": 100
    }
  }
}
```

## Environment Variables

The package supports the following environment variables, compatible with qase-java-commons:

### Main Configuration
- `QASE_MODE` - Operating mode: `testops`, `report`, `off`
- `QASE_FALLBACK` - Fallback mode: `report`, `off`
- `QASE_DEBUG` - Enable debug logging: `true`, `false`
- `QASE_ENVIRONMENT` - Environment name
- `QASE_CAPTURE_LOGS` - Capture logs: `true`, `false`

### Report Configuration
- `QASE_REPORT_DRIVER` - Report driver: `local`
- `QASE_REPORT_CONNECTION_PATH` - Report output path
- `QASE_REPORT_CONNECTION_FORMAT` - Report format: `json`, `yaml`

### TestOps API Configuration
- `QASE_TESTOPS_API_TOKEN` - API token for Qase TestOps
- `QASE_TESTOPS_API_HOST` - API host (default: `qase.io`)

### TestOps Run Configuration
- `QASE_TESTOPS_RUN_TITLE` - Test run title
- `QASE_TESTOPS_RUN_DESCRIPTION` - Test run description
- `QASE_TESTOPS_RUN_COMPLETE` - Auto-complete run: `true`, `false`
- `QASE_TESTOPS_RUN_TAGS` - Run tags (comma-separated)

### TestOps Other Configuration
- `QASE_TESTOPS_DEFECT` - Link defects: `true`, `false`
- `QASE_TESTOPS_PROJECT` - Project code
- `QASE_TESTOPS_BATCH_SIZE` - Batch size for sending results

## Configuration Priority

Configuration values are loaded with the following priority (highest to lowest):

1. **Environment Variables** (highest priority)
2. **Configuration File** (medium priority)
3. **Default Values** (lowest priority)

## Examples

### Create Default Configuration File

```go
package main

import (
    "log"
    
    "github.com/qase-tms/qase-go/config"
)

func main() {
    // Create default config file in current directory
    err := config.CreateDefaultConfig()
    if err != nil {
        log.Fatal(err)
    }
    
    // Or create at specific path
    err = config.CreateDefaultConfigAt("./configs/qase.config.json")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Advanced Configuration Loading

```go
package main

import (
    "log"
    
    "github.com/qase-tms/qase-go/config"
)

func main() {
    // Create custom loader
    loader := config.NewConfigLoader().
        WithSearchPaths([]string{
            "./config",
            "./configs",
            "./test/config",
            "/etc/qase",
        }).
        WithFileName("custom-qase-config.json")
    
    cfg, err := loader.Load()
    if err != nil {
        log.Fatal(err)
    }
    
    // Use configuration...
}
```

### Working with Configuration Values

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/qase-tms/qase-go/config"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }
    
    // Check mode
    switch cfg.Mode {
    case "testops":
        fmt.Println("TestOps mode enabled")
        if cfg.TestOps.API.Token == "" {
            log.Fatal("API token is required for TestOps mode")
        }
    case "report":
        fmt.Println("Report mode enabled")
        fmt.Printf("Report will be saved to: %s\n", cfg.Report.Connection.Local.Path)
    case "off":
        fmt.Println("Reporter is disabled")
        return
    }
    
    // Access nested configuration
    if cfg.Debug {
        fmt.Println("Debug logging enabled")
    }
    
    // Work with arrays
    if len(cfg.TestOps.Run.Tags) > 0 {
        fmt.Printf("Run tags: %v\n", cfg.TestOps.Run.Tags)
    }
    
    // Work with configurations
    for _, c := range cfg.TestOps.Run.Configurations.Values {
        fmt.Printf("Configuration: %s = %s\n", c.Name, c.Value)
    }
}
```

### Custom Defaults

```go
package main

import (
    "log"
    
    "github.com/qase-tms/qase-go/config"
)

func main() {
    // Create custom defaults
    defaults := config.NewConfig()
    defaults.Mode = "testops"
    defaults.Debug = true
    defaults.TestOps.Batch.Size = 200
    
    // Load with custom defaults
    loader := config.NewConfigLoader()
    cfg, err := loader.LoadWithDefaults(defaults)
    if err != nil {
        log.Fatal(err)
    }
    
    // Use configuration...
}
```

## API Reference

### Types

- `Config` - Main configuration structure
- `ConfigBuilder` - Builder for creating configurations
- `ConfigLoader` - Loader for loading configurations from multiple sources

### Functions

- `NewConfig() *Config` - Create new configuration with defaults
- `NewConfigBuilder() *ConfigBuilder` - Create new configuration builder
- `NewConfigLoader() *ConfigLoader` - Create new configuration loader
- `Load() (*Config, error)` - Load configuration using default loader
- `LoadFrom(filePath string) (*Config, error)` - Load configuration from specific file
- `CreateDefaultConfig() error` - Create default configuration file

### Configuration Modes

- `testops` - Send results to Qase TestOps
- `report` - Save results to local files
- `off` - Disable reporter

### Report Formats

- `json` - JSON format
- `yaml` - YAML format

## Testing

Run tests:

```bash
go test ./config
```

Run tests with coverage:

```bash
go test -cover ./config
```

## License

This package is part of the Qase Go SDK and follows the same license terms.