package clients

import (
	"strings"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

func TestV1Client_UserAgent(t *testing.T) {
	client, err := NewV1Client(ClientConfig{
		APIToken: "test-token",
	})
	if err != nil {
		t.Fatalf("failed to create V1Client: %v", err)
	}

	ua := client.GetAPIClient().GetConfig().UserAgent
	expected := "qase-api-client-go/" + domain.Version

	if ua != expected {
		t.Errorf("expected User-Agent %q, got %q", expected, ua)
	}

	if !strings.Contains(ua, "qase-api-client") {
		t.Errorf("User-Agent must contain 'qase-api-client', got %q", ua)
	}
}

func TestV2Client_UserAgent(t *testing.T) {
	client, err := NewV2Client(ClientConfig{
		APIToken: "test-token",
	})
	if err != nil {
		t.Fatalf("failed to create V2Client: %v", err)
	}

	ua := client.GetAPIClient().GetConfig().UserAgent
	expected := "qase-api-client-go/" + domain.Version

	if ua != expected {
		t.Errorf("expected User-Agent %q, got %q", expected, ua)
	}

	if !strings.Contains(ua, "qase-api-client") {
		t.Errorf("User-Agent must contain 'qase-api-client', got %q", ua)
	}
}

func TestUserAgent_Format(t *testing.T) {
	client, err := NewV1Client(ClientConfig{
		APIToken: "test-token",
	})
	if err != nil {
		t.Fatalf("failed to create V1Client: %v", err)
	}

	ua := client.GetAPIClient().GetConfig().UserAgent

	// Must match format: qase-api-client-<language>/<version>
	if !strings.HasPrefix(ua, "qase-api-client-go/") {
		t.Errorf("User-Agent must start with 'qase-api-client-go/', got %q", ua)
	}

	parts := strings.SplitN(ua, "/", 2)
	if len(parts) != 2 {
		t.Fatalf("User-Agent must contain '/', got %q", ua)
	}

	version := parts[1]
	if version == "" {
		t.Error("User-Agent version part must not be empty")
	}

	if version != domain.Version {
		t.Errorf("User-Agent version must match domain.Version %q, got %q", domain.Version, version)
	}
}
