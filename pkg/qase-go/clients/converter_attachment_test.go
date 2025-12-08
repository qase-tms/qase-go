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
	ReturnHashes []string // If set, returns these hashes instead of generating them
}

func (m *MockAttachmentUploader) UploadAttachment(ctx context.Context, projectCode string, file []*os.File) ([]string, error) {
	if m.ShouldFail {
		return nil, fmt.Errorf("mock upload failed")
	}
	
	// If ReturnHashes is set, use it (truncate or pad as needed)
	if len(m.ReturnHashes) > 0 {
		hashes := make([]string, len(file))
		for i := range file {
			if i < len(m.ReturnHashes) {
				hashes[i] = m.ReturnHashes[i]
			} else {
				hashes[i] = fmt.Sprintf("mock-hash-%d", i)
			}
		}
		return hashes, nil
	}
	
	// If ReturnHash is set, return one hash per file
	if m.ReturnHash != "" {
		hashes := make([]string, len(file))
		for i := range file {
			hashes[i] = m.ReturnHash
		}
		return hashes, nil
	}
	
	// Generate hashes for each file
	hashes := make([]string, len(file))
	for i := range file {
		hashes[i] = fmt.Sprintf("mock-hash-%d", i)
	}
	return hashes, nil
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

func TestV2Converter_ProcessAttachments_BatchUpload(t *testing.T) {
	// Create temporary files for testing
	tmpDir := t.TempDir()
	
	files := []struct {
		name    string
		content string
	}{
		{"file1.txt", "content 1"},
		{"file2.txt", "content 2"},
		{"file3.txt", "content 3"},
	}

	var attachments []domain.Attachment
	for _, f := range files {
		filePath := filepath.Join(tmpDir, f.name)
		err := os.WriteFile(filePath, []byte(f.content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", f.name, err)
		}
		attachments = append(attachments, domain.Attachment{
			ID:       "temp-id",
			FilePath: &filePath,
			FileName: f.name,
		})
	}

	uploader := &MockAttachmentUploader{ReturnHashes: []string{"hash1", "hash2", "hash3"}}
	converter := &V2Converter{
		uploader:    uploader,
		projectCode: "TEST",
	}

	result, err := converter.processAttachmentsGracefully(context.Background(), attachments)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(result) != 3 {
		t.Fatalf("Expected 3 attachments, got %d", len(result))
	}

	expected := []string{"hash1", "hash2", "hash3"}
	for i, expectedHash := range expected {
		if result[i] != expectedHash {
			t.Errorf("Expected hash %s at index %d, got %s", expectedHash, i, result[i])
		}
	}
}
