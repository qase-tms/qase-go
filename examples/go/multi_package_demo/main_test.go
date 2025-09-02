package main

import (
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

// TestMain demonstrates global initialization for multi-package test structure
func TestMain(m *testing.M) {
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

func TestMainPackageTest(t *testing.T) {
	qase.Test(t,
		qase.TestMetadata{
			DisplayName: "Main Package Test",
			Description: "A test in the main package that uses global initialization",
		},
		func() {
			qase.AddMessage("Starting main package test")
			qase.True(t, true)
			qase.Equal(t, 1+1, 2)
			qase.AddMessage("Main package test completed successfully")
		})
}
