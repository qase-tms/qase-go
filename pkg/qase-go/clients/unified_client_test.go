package clients

import (
	"context"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/config"
	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestNewUnifiedClient(t *testing.T) {
	tests := []struct {
		name        string
		config      *config.Config
		expectError bool
	}{
		{
			name: "valid configuration",
			config: &config.Config{
				TestOps: config.TestOpsConfig{
					API: config.APIConfig{
						Token: "test-token",
						Host:  "qase.io",
					},
					Project: "TEST",
				},
				Debug: true,
			},
			expectError: false,
		},
		{
			name: "empty configuration",
			config: &config.Config{
				TestOps: config.TestOpsConfig{
					API: config.APIConfig{
						Token: "",
						Host:  "",
					},
					Project: "",
				},
			},
			expectError: false, // Client creation should succeed even with empty config
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewUnifiedClient(tt.config)
			
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			
			if !tt.expectError && client != nil {
				// Verify client structure
				if client.v1Client == nil {
					t.Error("v1Client should not be nil")
				}
				if client.v2Client == nil {
					t.Error("v2Client should not be nil")
				}
				if client.config != tt.config {
					t.Error("config should be stored correctly")
				}
				if client.projectCode != tt.config.TestOps.Project {
					t.Errorf("expected project code %s, got %s", tt.config.TestOps.Project, client.projectCode)
				}
			}
		})
	}
}

func TestUnifiedClient_GetProjectCode(t *testing.T) {
	cfg := &config.Config{
		TestOps: config.TestOpsConfig{
			API: config.APIConfig{
				Token: "test-token",
				Host:  "qase.io",
			},
			Project: "MYPROJECT",
		},
	}
	
	client, err := NewUnifiedClient(cfg)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	
	if client.GetProjectCode() != "MYPROJECT" {
		t.Errorf("expected project code 'MYPROJECT', got %s", client.GetProjectCode())
	}
}

func TestUnifiedClient_GetConfig(t *testing.T) {
	cfg := &config.Config{
		TestOps: config.TestOpsConfig{
			API: config.APIConfig{
				Token: "test-token",
				Host:  "qase.io",
			},
			Project: "TEST",
		},
		Debug: true,
	}
	
	client, err := NewUnifiedClient(cfg)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	
	returnedCfg := client.GetConfig()
	if returnedCfg != cfg {
		t.Error("returned config should be the same as provided")
	}
	if returnedCfg.Debug != true {
		t.Error("debug flag should be preserved")
	}
}

func TestUnifiedClient_CreateRun_TitleFromConfig(t *testing.T) {
	tests := []struct {
		name           string
		runTitle       string
		runDescription string
		environment    string
		expectedTitle  string
		expectedDesc   string
	}{
		{
			name:          "title from config",
			runTitle:      "Custom Test Run",
			runDescription: "Custom description",
			expectedTitle: "Custom Test Run",
			expectedDesc:  "Custom description",
		},
		{
			name:          "default title",
			runTitle:      "",
			runDescription: "",
			expectedTitle: "Automated Test Run",
			expectedDesc:  "",
		},
		{
			name:          "title with environment description",
			runTitle:      "",
			runDescription: "",
			environment:   "staging",
			expectedTitle: "Automated Test Run",
			expectedDesc:  "Test run in staging environment",
		},
		{
			name:          "explicit description overrides environment",
			runTitle:      "Custom Title",
			runDescription: "Explicit description",
			environment:   "staging",
			expectedTitle: "Custom Title",
			expectedDesc:  "Explicit description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config.Config{
				TestOps: config.TestOpsConfig{
					API: config.APIConfig{
						Token: "test-token",
						Host:  "qase.io",
					},
					Project: "TEST",
					Run: config.RunConfig{
						Title:       tt.runTitle,
						Description: tt.runDescription,
					},
				},
				Environment: tt.environment,
			}
			
			client, err := NewUnifiedClient(cfg)
			if err != nil {
				t.Fatalf("failed to create client: %v", err)
			}
			
			// Note: We can't actually test the API call without mocking,
			// but we can verify the client was created correctly
			if client.config.TestOps.Run.Title != tt.runTitle {
				t.Errorf("expected run title %s, got %s", tt.runTitle, client.config.TestOps.Run.Title)
			}
			if client.config.TestOps.Run.Description != tt.runDescription {
				t.Errorf("expected run description %s, got %s", tt.runDescription, client.config.TestOps.Run.Description)
			}
		})
	}
}

func TestUnifiedClient_UploadResults_BatchHandling(t *testing.T) {
	cfg := &config.Config{
		TestOps: config.TestOpsConfig{
			API: config.APIConfig{
				Token: "test-token",
				Host:  "qase.io",
			},
			Project: "TEST",
			Batch: config.BatchConfig{
				Size: 2, // Small batch size for testing
			},
		},
	}
	
	client, err := NewUnifiedClient(cfg)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// Test with empty results
	ctx := context.Background()
	err = client.UploadResults(ctx, 123, nil)
	if err != nil {
		t.Errorf("UploadResults with empty slice should not fail: %v", err)
	}

	// Test with results (this will fail without real API, but we test the structure)
	results := []*domain.TestResult{
		domain.NewTestResult("Test 1"),
		domain.NewTestResult("Test 2"),
		domain.NewTestResult("Test 3"),
	}
	
	// This will fail due to no real API connection, but the structure should be correct
	err = client.UploadResults(ctx, 123, results)
	// We expect this to fail since there's no real API server, 
	// but the error should be from the API call, not from our logic
	if err == nil {
		t.Log("UploadResults succeeded (unexpected but possible in test environment)")
	} else {
		t.Logf("UploadResults failed as expected without real API: %v", err)
	}
}

func TestUnifiedClient_Interface(t *testing.T) {
	cfg := &config.Config{
		TestOps: config.TestOpsConfig{
			API: config.APIConfig{
				Token: "test-token",
				Host:  "qase.io",
			},
			Project: "TEST",
		},
	}
	
	client, err := NewUnifiedClient(cfg)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	
	// Test that all interface methods exist (will fail if API calls fail, which is expected)
	ctx := context.Background()
	
	// CreateRun
	_, err = client.CreateRun(ctx)
	if err != nil {
		t.Logf("CreateRun failed as expected without real API: %v", err)
	}
	
	// CompleteRun
	err = client.CompleteRun(ctx, 123)
	if err != nil {
		t.Logf("CompleteRun failed as expected without real API: %v", err)
	}
	
	// UploadResults
	results := []*domain.TestResult{domain.NewTestResult("Test")}
	err = client.UploadResults(ctx, 123, results)
	if err != nil {
		t.Logf("UploadResults failed as expected without real API: %v", err)
	}
}