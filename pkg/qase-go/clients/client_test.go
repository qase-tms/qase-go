package clients

import "testing"

func TestClients_ResolveServerURLsFromHost(t *testing.T) {
	tests := []struct {
		name   string
		host   string
		wantV1 string
		wantV2 string
	}{
		{"empty host uses default", "", "https://api.qase.io/v1", "https://api.qase.io/v2"},
		{"qase.io uses default", "qase.io", "https://api.qase.io/v1", "https://api.qase.io/v2"},
		{"EU cluster", "eu.qase.io", "https://api-eu.qase.io/v1", "https://api-eu.qase.io/v2"},
		{"enterprise host", "example.qase.io", "https://api-example.qase.io/v1", "https://api-example.qase.io/v2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientConfig := ClientConfig{BaseURL: buildAPIBaseURL(tt.host), APIToken: "token"}

			v1Client, err := NewV1Client(clientConfig)
			if err != nil {
				t.Fatalf("NewV1Client: %v", err)
			}
			v2Client, err := NewV2Client(clientConfig)
			if err != nil {
				t.Fatalf("NewV2Client: %v", err)
			}

			v1Servers := v1Client.client.GetConfig().Servers
			if len(v1Servers) != 1 || v1Servers[0].URL != tt.wantV1 {
				t.Errorf("v1 servers = %+v, want single %q", v1Servers, tt.wantV1)
			}
			v2Servers := v2Client.client.GetConfig().Servers
			if len(v2Servers) != 1 || v2Servers[0].URL != tt.wantV2 {
				t.Errorf("v2 servers = %+v, want single %q", v2Servers, tt.wantV2)
			}
		})
	}
}
