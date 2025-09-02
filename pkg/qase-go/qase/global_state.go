package qase

import (
	"sync"
	"sync/atomic"

	"github.com/qase-tms/qase-go/pkg/qase-go/reporters"
)

// Global state shared across all packages
var (
	globalInitialized int32
	globalReporter    *reporters.CoreReporter
	globalMutex       sync.RWMutex
)

// SetGlobalInitialized sets the global initialization state
func SetGlobalInitialized(initialized bool) {
	if initialized {
		atomic.StoreInt32(&globalInitialized, 1)
	} else {
		atomic.StoreInt32(&globalInitialized, 0)
	}
}

// IsGlobalInitialized returns true if qase has been globally initialized
func IsGlobalInitialized() bool {
	return atomic.LoadInt32(&globalInitialized) == 1
}

// SetGlobalReporter sets the global reporter instance
func SetGlobalReporter(reporter *reporters.CoreReporter) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	globalReporter = reporter
}

// GetGlobalReporter returns the global reporter instance if available
func GetGlobalReporter() *reporters.CoreReporter {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	return globalReporter
}

// ResetGlobalState resets the global state (useful for testing)
func ResetGlobalState() {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	atomic.StoreInt32(&globalInitialized, 0)
	globalReporter = nil
}
