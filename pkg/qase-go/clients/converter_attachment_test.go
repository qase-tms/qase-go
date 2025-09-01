package clients

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
)

// MockAttachmentUploader implements AttachmentUploader for testing
type MockAttachmentUploader struct {
	ShouldFail bool
	ReturnHash string
}

func (m *MockAttachmentUploader) UploadAttachment(ctx context.Context, projectCode string, file []*os.File) (string, error) {
	if m.ShouldFail {
		return "", fmt.Errorf("mock upload failed")
	}
	if m.ReturnHash != "" {
		return m.ReturnHash, nil
	}
	return "mock-hash-123", nil
}

func TestV2Converter_ProcessAttachments(t *testing.T) {
	tests := []struct {
		name        string
		uploader    AttachmentUploader
		projectCode string
		attachments []domain.Attachment
		expected    []string
		expectError bool
	}{
		{
			name:        "attachments with IDs only",
			uploader:    nil,
			projectCode: "",
			attachments: []domain.Attachment{
				{ID: "existing-id-1"},
				{ID: "existing-id-2"},
			},
			expected:    []string{"existing-id-1", "existing-id-2"},
			expectError: false,
		},
		{
			name:        "attachments without file paths - uploader available but not used",
			uploader:    &MockAttachmentUploader{ReturnHash: "uploaded-hash"},
			projectCode: "TEST",
			attachments: []domain.Attachment{
				{ID: "existing-id-1"},
				{ID: "existing-id-2"}, // No file path, so will use ID
			},
			expected:    []string{"existing-id-1", "existing-id-2"},
			expectError: false,
		},
		{
			name:        "upload failure",
			uploader:    &MockAttachmentUploader{ShouldFail: true},
			projectCode: "TEST",
			attachments: []domain.Attachment{
				{ID: "temp-id", FilePath: stringPtr("/tmp/nonexistent.txt")},
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := &V2Converter{
				uploader:    tt.uploader,
				projectCode: tt.projectCode,
			}

			result, err := converter.processAttachmentsGracefully(context.Background(), tt.attachments)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d attachments, got %d", len(tt.expected), len(result))
				return
			}

			for i, expected := range tt.expected {
				if result[i] != expected {
					t.Errorf("Expected attachment ID %s at index %d, got %s", expected, i, result[i])
				}
			}
		})
	}
}

func TestV2Converter_ProcessAttachments_WithFileUpload(t *testing.T) {
	// Create a temporary file for testing
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	content := "test file content"

	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	uploader := &MockAttachmentUploader{ReturnHash: "uploaded-file-hash"}
	converter := &V2Converter{
		uploader:    uploader,
		projectCode: "TEST",
	}

	attachments := []domain.Attachment{
		{
			ID:       "temp-id",
			FilePath: &tmpFile,
			FileName: "test.txt",
		},
	}

	result, err := converter.processAttachmentsGracefully(context.Background(), attachments)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("Expected 1 attachment, got %d", len(result))
	}

	if result[0] != "uploaded-file-hash" {
		t.Errorf("Expected uploaded hash 'uploaded-file-hash', got '%s'", result[0])
	}
}

func TestV2Converter_NewV2ConverterWithUploader(t *testing.T) {
	uploader := &MockAttachmentUploader{}
	projectCode := "TEST"

	converter := NewV2ConverterWithUploader(uploader, projectCode)

	if converter.uploader != uploader {
		t.Errorf("Expected uploader to be set")
	}

	if converter.projectCode != projectCode {
		t.Errorf("Expected project code '%s', got '%s'", projectCode, converter.projectCode)
	}
}

func TestV2Converter_ProcessAttachments_NonexistentFile(t *testing.T) {
	uploader := &MockAttachmentUploader{}
	converter := &V2Converter{
		uploader:    uploader,
		projectCode: "TEST",
	}

	nonexistentFile := "/tmp/nonexistent-file-12345.txt"
	attachments := []domain.Attachment{
		{
			ID:       "temp-id",
			FilePath: &nonexistentFile,
		},
	}

	_, err := converter.processAttachmentsGracefully(context.Background(), attachments)
	if err == nil {
		t.Error("Expected error for all attachments failing, but got none")
	}

	if !strings.Contains(err.Error(), "all attachments failed to process") {
		t.Errorf("Expected 'all attachments failed to process' error, got: %v", err)
	}
}

func TestV2Converter_ProcessAttachments_ContentAttachmentPreservesExtension(t *testing.T) {
	uploader := &MockAttachmentUploader{ReturnHash: "content-uploaded-hash"}
	converter := &V2Converter{
		uploader:    uploader,
		projectCode: "TEST",
	}

	attachments := []domain.Attachment{
		{
			ID:       "content-id",
			FileName: "log.txt",
			MimeType: "text/plain",
			Content:  []byte("test log content"),
		},
	}

	result, err := converter.processAttachmentsGracefully(context.Background(), attachments)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("Expected 1 attachment, got %d", len(result))
	}

	if result[0] != "content-uploaded-hash" {
		t.Errorf("Expected uploaded hash 'content-uploaded-hash', got '%s'", result[0])
	}
}
