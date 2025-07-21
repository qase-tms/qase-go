package reporters

import (
	"fmt"
	"sync"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
)

// ReporterFactory manages the singleton instance of CoreReporter
type ReporterFactory struct {
	instance *CoreReporter
	mutex    sync.RWMutex
	config   *config.Config
}

// NewReporterFactory creates a new reporter factory
func NewReporterFactory(cfg *config.Config) *ReporterFactory {
	return &ReporterFactory{
		config: cfg,
	}
}

// GetReporter returns the singleton instance of CoreReporter
func (rf *ReporterFactory) GetReporter() (*CoreReporter, error) {
	rf.mutex.RLock()
	if rf.instance != nil {
		defer rf.mutex.RUnlock()
		return rf.instance, nil
	}
	rf.mutex.RUnlock()

	// Double-checked locking pattern
	rf.mutex.Lock()
	defer rf.mutex.Unlock()

	if rf.instance != nil {
		return rf.instance, nil
	}

	// Create new instance
	reporter, err := NewCoreReporter(rf.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create core reporter: %w", err)
	}

	rf.instance = reporter
	return rf.instance, nil
}

// Reset clears the singleton instance (useful for testing)
func (rf *ReporterFactory) Reset() {
	rf.mutex.Lock()
	defer rf.mutex.Unlock()
	rf.instance = nil
}

// UpdateConfig updates the configuration and resets the instance
func (rf *ReporterFactory) UpdateConfig(cfg *config.Config) {
	rf.mutex.Lock()
	defer rf.mutex.Unlock()
	rf.config = cfg
	rf.instance = nil
}

// GetConfig returns the current configuration
func (rf *ReporterFactory) GetConfig() *config.Config {
	rf.mutex.RLock()
	defer rf.mutex.RUnlock()
	return rf.config
}

// IsInitialized returns true if the reporter has been initialized
func (rf *ReporterFactory) IsInitialized() bool {
	rf.mutex.RLock()
	defer rf.mutex.RUnlock()
	return rf.instance != nil
}

// Global factory instance
var (
	globalFactory *ReporterFactory
	globalMutex   sync.RWMutex
)

// InitializeGlobalFactory initializes the global factory with the given configuration
func InitializeGlobalFactory(cfg *config.Config) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	globalFactory = NewReporterFactory(cfg)
}

// GetGlobalReporter returns the global singleton instance of CoreReporter
func GetGlobalReporter() (*CoreReporter, error) {
	globalMutex.RLock()
	if globalFactory == nil {
		globalMutex.RUnlock()
		return nil, fmt.Errorf("global factory not initialized - call InitializeGlobalFactory first")
	}
	globalMutex.RUnlock()

	return globalFactory.GetReporter()
}

// ResetGlobalFactory resets the global factory (useful for testing)
func ResetGlobalFactory() {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	globalFactory = nil
}

// UpdateGlobalConfig updates the global configuration
func UpdateGlobalConfig(cfg *config.Config) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	if globalFactory == nil {
		globalFactory = NewReporterFactory(cfg)
	} else {
		globalFactory.UpdateConfig(cfg)
	}
}

// IsGlobalFactoryInitialized returns true if the global factory has been initialized
func IsGlobalFactoryInitialized() bool {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	return globalFactory != nil
}
