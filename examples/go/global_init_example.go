package _go

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

// TestMain demonstrates global initialization for multi-package test structure
// This should be placed in the root package of your test suite
func TestMainGlobalInit(m *testing.M) {
	// Initialize qase globally - this will search for config in parent directories
	// and initialize the reporter once for all packages
	if err := qase.InitializeGlobal(); err != nil {
		// If global initialization fails, you can still run tests
		// but they won't be reported to qase
		panic("Failed to initialize qase globally: " + err.Error())
	}

	// Run all tests
	_ = m.Run()

	// Complete test run automatically after all tests finish
	qase.TestMain(m)
}
