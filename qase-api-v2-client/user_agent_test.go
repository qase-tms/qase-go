package api_v2_client

import (
	"strings"
	"testing"
)

func TestNewConfiguration_UserAgent(t *testing.T) {
	cfg := NewConfiguration()

	if !strings.Contains(cfg.UserAgent, "qase-api-client") {
		t.Errorf("User-Agent must contain 'qase-api-client', got %q", cfg.UserAgent)
	}

	if !strings.HasPrefix(cfg.UserAgent, "qase-api-client-go/") {
		t.Errorf("User-Agent must start with 'qase-api-client-go/', got %q", cfg.UserAgent)
	}

	parts := strings.SplitN(cfg.UserAgent, "/", 2)
	if len(parts) != 2 || parts[1] == "" {
		t.Errorf("User-Agent must have format 'qase-api-client-go/<version>', got %q", cfg.UserAgent)
	}
}

func TestGetModuleVersion(t *testing.T) {
	version := getModuleVersion()
	if version == "" {
		t.Error("module version must not be empty")
	}

	// Version should not contain "v" prefix
	if strings.HasPrefix(version, "v") {
		t.Errorf("module version should not have 'v' prefix, got %q", version)
	}
}
