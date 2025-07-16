package config

// ConfigBuilder provides a fluent interface for building configuration
type ConfigBuilder struct {
	config *Config
}

// NewConfigBuilder creates a new ConfigBuilder with default configuration
func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: NewConfig(),
	}
}

// FromFile loads configuration from file and returns builder
func FromFile(filename string) (*ConfigBuilder, error) {
	config, err := LoadFromFile(filename)
	if err != nil {
		return nil, err
	}
	return &ConfigBuilder{config: config}, nil
}

// Main configuration methods

// WithMode sets the operating mode
func (b *ConfigBuilder) WithMode(mode string) *ConfigBuilder {
	b.config.Mode = mode
	return b
}

// WithFallback sets the fallback mode
func (b *ConfigBuilder) WithFallback(fallback string) *ConfigBuilder {
	b.config.Fallback = fallback
	return b
}

// WithDebug sets debug flag
func (b *ConfigBuilder) WithDebug(debug bool) *ConfigBuilder {
	b.config.Debug = debug
	return b
}

// WithEnvironment sets environment name
func (b *ConfigBuilder) WithEnvironment(environment string) *ConfigBuilder {
	b.config.Environment = environment
	return b
}

// WithCaptureLogs sets capture logs flag
func (b *ConfigBuilder) WithCaptureLogs(captureLogs bool) *ConfigBuilder {
	b.config.CaptureLogs = captureLogs
	return b
}

// Report configuration methods

// WithReportDriver sets report driver
func (b *ConfigBuilder) WithReportDriver(driver string) *ConfigBuilder {
	b.config.Report.Driver = driver
	return b
}

// WithReportPath sets report output path
func (b *ConfigBuilder) WithReportPath(path string) *ConfigBuilder {
	b.config.Report.Connection.Local.Path = path
	return b
}

// WithReportFormat sets report format
func (b *ConfigBuilder) WithReportFormat(format string) *ConfigBuilder {
	b.config.Report.Connection.Local.Format = format
	return b
}

// TestOps API configuration methods

// WithAPIToken sets API token
func (b *ConfigBuilder) WithAPIToken(token string) *ConfigBuilder {
	b.config.TestOps.API.Token = token
	return b
}

// WithAPIHost sets API host
func (b *ConfigBuilder) WithAPIHost(host string) *ConfigBuilder {
	b.config.TestOps.API.Host = host
	return b
}

// TestOps Run configuration methods

// WithRunTitle sets run title
func (b *ConfigBuilder) WithRunTitle(title string) *ConfigBuilder {
	b.config.TestOps.Run.Title = title
	return b
}

// WithRunDescription sets run description
func (b *ConfigBuilder) WithRunDescription(description string) *ConfigBuilder {
	b.config.TestOps.Run.Description = description
	return b
}

// WithRunComplete sets run complete flag
func (b *ConfigBuilder) WithRunComplete(complete bool) *ConfigBuilder {
	b.config.TestOps.Run.Complete = complete
	return b
}

// WithRunTags sets run tags
func (b *ConfigBuilder) WithRunTags(tags []string) *ConfigBuilder {
	b.config.TestOps.Run.Tags = tags
	return b
}

// AddRunTag adds a single tag to run tags
func (b *ConfigBuilder) AddRunTag(tag string) *ConfigBuilder {
	b.config.TestOps.Run.Tags = append(b.config.TestOps.Run.Tags, tag)
	return b
}

// WithRunConfigurations sets run configurations
func (b *ConfigBuilder) WithRunConfigurations(configurations []ConfigurationValue) *ConfigBuilder {
	b.config.TestOps.Run.Configurations.Values = configurations
	return b
}

// AddRunConfiguration adds a single configuration to run configurations
func (b *ConfigBuilder) AddRunConfiguration(name, value string) *ConfigBuilder {
	b.config.TestOps.Run.Configurations.Values = append(
		b.config.TestOps.Run.Configurations.Values,
		ConfigurationValue{Name: name, Value: value},
	)
	return b
}

// WithCreateConfigurationsIfNotExists sets flag to create configurations if not exists
func (b *ConfigBuilder) WithCreateConfigurationsIfNotExists(create bool) *ConfigBuilder {
	b.config.TestOps.Run.Configurations.CreateIfNotExists = create
	return b
}

// TestOps other configuration methods

// WithDefect sets defect flag
func (b *ConfigBuilder) WithDefect(defect bool) *ConfigBuilder {
	b.config.TestOps.Defect = defect
	return b
}

// WithProject sets project code
func (b *ConfigBuilder) WithProject(project string) *ConfigBuilder {
	b.config.TestOps.Project = project
	return b
}

// WithBatchSize sets batch size
func (b *ConfigBuilder) WithBatchSize(size int) *ConfigBuilder {
	b.config.TestOps.Batch.Size = size
	return b
}

// Environment and file loading methods

// LoadFromEnvironment loads configuration from environment variables
func (b *ConfigBuilder) LoadFromEnvironment() *ConfigBuilder {
	b.config.LoadFromEnvironment()
	return b
}

// LoadFromFileIfExists loads configuration from file if it exists
func (b *ConfigBuilder) LoadFromFileIfExists(filename string) *ConfigBuilder {
	if config, err := LoadFromFile(filename); err == nil {
		b.config = config
	}
	return b
}

// Build returns the built configuration after validation
func (b *ConfigBuilder) Build() (*Config, error) {
	if err := b.config.Validate(); err != nil {
		return nil, err
	}
	return b.config, nil
}

// BuildUnsafe returns the built configuration without validation
func (b *ConfigBuilder) BuildUnsafe() *Config {
	return b.config
}

// Preset methods for common configurations

// TestOpsMode configures builder for TestOps mode
func (b *ConfigBuilder) TestOpsMode() *ConfigBuilder {
	return b.WithMode("testops")
}

// ReportMode configures builder for Report mode
func (b *ConfigBuilder) ReportMode() *ConfigBuilder {
	return b.WithMode("report")
}

// OffMode configures builder for Off mode
func (b *ConfigBuilder) OffMode() *ConfigBuilder {
	return b.WithMode("off")
}

// JSONReport configures builder for JSON report format
func (b *ConfigBuilder) JSONReport() *ConfigBuilder {
	return b.WithReportFormat("json")
}

// YAMLReport configures builder for YAML report format
func (b *ConfigBuilder) YAMLReport() *ConfigBuilder {
	return b.WithReportFormat("yaml")
}