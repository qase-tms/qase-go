package domain

import (
	"encoding/json"
	"testing"
)

func TestNewAttachment(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		fileName string
		mimeType string
		content  []byte
		expected *Attachment
	}{
		{
			name:     "basic attachment",
			id:       "att-123",
			fileName: "test.txt",
			mimeType: "text/plain",
			content:  []byte("test content"),
			expected: &Attachment{
				ID:       "att-123",
				FileName: "test.txt",
				MimeType: "text/plain",
				Content:  []byte("test content"),
				Size:     12,
				FilePath: nil,
			},
		},
		{
			name:     "empty content",
			id:       "att-empty",
			fileName: "empty.txt",
			mimeType: "text/plain",
			content:  []byte{},
			expected: &Attachment{
				ID:       "att-empty",
				FileName: "empty.txt",
				MimeType: "text/plain",
				Content:  []byte{},
				Size:     0,
				FilePath: nil,
			},
		},
		{
			name:     "binary content",
			id:       "att-bin",
			fileName: "image.jpg",
			mimeType: "image/jpeg",
			content:  []byte{0xFF, 0xD8, 0xFF, 0xE0},
			expected: &Attachment{
				ID:       "att-bin",
				FileName: "image.jpg",
				MimeType: "image/jpeg",
				Content:  []byte{0xFF, 0xD8, 0xFF, 0xE0},
				Size:     4,
				FilePath: nil,
			},
		},
		{
			name:     "large content",
			id:       "att-large",
			fileName: "large.dat",
			mimeType: "application/octet-stream",
			content:  make([]byte, 1024*1024), // 1MB
			expected: &Attachment{
				ID:       "att-large",
				FileName: "large.dat",
				MimeType: "application/octet-stream",
				Content:  make([]byte, 1024*1024),
				Size:     1024 * 1024,
				FilePath: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewAttachment(tt.id, tt.fileName, tt.mimeType, tt.content)

			if result.ID != tt.expected.ID {
				t.Errorf("ID: got %q, want %q", result.ID, tt.expected.ID)
			}
			if result.FileName != tt.expected.FileName {
				t.Errorf("FileName: got %q, want %q", result.FileName, tt.expected.FileName)
			}
			if result.MimeType != tt.expected.MimeType {
				t.Errorf("MimeType: got %q, want %q", result.MimeType, tt.expected.MimeType)
			}
			if result.Size != tt.expected.Size {
				t.Errorf("Size: got %d, want %d", result.Size, tt.expected.Size)
			}
			if result.FilePath != nil {
				t.Errorf("FilePath: got %v, want nil", result.FilePath)
			}
			if len(result.Content) != len(tt.expected.Content) {
				t.Errorf("Content length: got %d, want %d", len(result.Content), len(tt.expected.Content))
			}
		})
	}
}

func TestAttachment_SetFilePath(t *testing.T) {
	attachment := NewAttachment("att-1", "test.txt", "text/plain", []byte("content"))

	tests := []struct {
		name string
		path string
	}{
		{
			name: "absolute path",
			path: "/path/to/file.txt",
		},
		{
			name: "relative path",
			path: "./relative/path.txt",
		},
		{
			name: "empty path",
			path: "",
		},
		{
			name: "path with spaces",
			path: "/path with spaces/file.txt",
		},
		{
			name: "windows path",
			path: "C:\\Users\\test\\file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attachment.SetFilePath(tt.path)

			if attachment.FilePath == nil {
				t.Error("FilePath should not be nil after SetFilePath")
				return
			}

			if *attachment.FilePath != tt.path {
				t.Errorf("FilePath: got %q, want %q", *attachment.FilePath, tt.path)
			}

			if !attachment.HasFilePath() {
				t.Error("HasFilePath should return true after SetFilePath")
			}

			if attachment.GetFilePath() != tt.path {
				t.Errorf("GetFilePath: got %q, want %q", attachment.GetFilePath(), tt.path)
			}
		})
	}
}

func TestAttachment_GetFilePath(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Attachment)
		expected string
	}{
		{
			name: "with file path set",
			setup: func(a *Attachment) {
				a.SetFilePath("/test/path.txt")
			},
			expected: "/test/path.txt",
		},
		{
			name:     "without file path",
			setup:    func(a *Attachment) {},
			expected: "",
		},
		{
			name: "empty file path",
			setup: func(a *Attachment) {
				a.SetFilePath("")
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attachment := NewAttachment("att-1", "test.txt", "text/plain", []byte("content"))
			tt.setup(attachment)

			result := attachment.GetFilePath()
			if result != tt.expected {
				t.Errorf("GetFilePath: got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestAttachment_HasFilePath(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Attachment)
		expected bool
	}{
		{
			name: "with file path set",
			setup: func(a *Attachment) {
				a.SetFilePath("/test/path.txt")
			},
			expected: true,
		},
		{
			name:     "without file path",
			setup:    func(a *Attachment) {},
			expected: false,
		},
		{
			name: "with empty file path",
			setup: func(a *Attachment) {
				a.SetFilePath("")
			},
			expected: true, // empty string is still a set value
		},
		{
			name: "manually set to nil",
			setup: func(a *Attachment) {
				a.SetFilePath("/test")
				a.FilePath = nil
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attachment := NewAttachment("att-1", "test.txt", "text/plain", []byte("content"))
			tt.setup(attachment)

			result := attachment.HasFilePath()
			if result != tt.expected {
				t.Errorf("HasFilePath: got %t, want %t", result, tt.expected)
			}
		})
	}
}

func TestAttachment_JSONSerialization(t *testing.T) {
	tests := []struct {
		name       string
		attachment *Attachment
		expected   string
	}{
		{
			name: "attachment without file path",
			attachment: &Attachment{
				ID:       "att-123",
				FileName: "test.txt",
				MimeType: "text/plain",
				Content:  []byte("hello"),
				Size:     5,
				FilePath: nil,
			},
			expected: `{"id":"att-123","file_name":"test.txt","mime_type":"text/plain","file_path":null,"content":"aGVsbG8=","size":5}`,
		},
		{
			name: "attachment with file path",
			attachment: func() *Attachment {
				a := NewAttachment("att-456", "image.jpg", "image/jpeg", []byte{0xFF, 0xD8})
				a.SetFilePath("/uploads/image.jpg")
				return a
			}(),
			expected: `{"id":"att-456","file_name":"image.jpg","mime_type":"image/jpeg","file_path":"/uploads/image.jpg","content":"/9g=","size":2}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.attachment)
			if err != nil {
				t.Fatalf("Failed to marshal attachment: %v", err)
			}

			if string(data) != tt.expected {
				t.Errorf("JSON serialization:\ngot:  %s\nwant: %s", string(data), tt.expected)
			}

			// Test deserialization
			var unmarshaled Attachment
			if err := json.Unmarshal(data, &unmarshaled); err != nil {
				t.Fatalf("Failed to unmarshal attachment: %v", err)
			}

			if unmarshaled.ID != tt.attachment.ID {
				t.Errorf("Unmarshaled ID: got %q, want %q", unmarshaled.ID, tt.attachment.ID)
			}
			if unmarshaled.FileName != tt.attachment.FileName {
				t.Errorf("Unmarshaled FileName: got %q, want %q", unmarshaled.FileName, tt.attachment.FileName)
			}
			if unmarshaled.MimeType != tt.attachment.MimeType {
				t.Errorf("Unmarshaled MimeType: got %q, want %q", unmarshaled.MimeType, tt.attachment.MimeType)
			}
			if unmarshaled.Size != tt.attachment.Size {
				t.Errorf("Unmarshaled Size: got %d, want %d", unmarshaled.Size, tt.attachment.Size)
			}
		})
	}
}

func TestAttachment_EdgeCases(t *testing.T) {
	t.Run("nil content", func(t *testing.T) {
		attachment := NewAttachment("att-1", "test.txt", "text/plain", nil)
		if attachment.Size != 0 {
			t.Errorf("Size with nil content: got %d, want 0", attachment.Size)
		}
		// Note: Go allows nil slices, and len(nil) == 0, so this is valid behavior
		if attachment.Content != nil && len(attachment.Content) != 0 {
			t.Error("Content should be nil or empty")
		}
	})

	t.Run("unicode filename", func(t *testing.T) {
		attachment := NewAttachment("att-unicode", "测试.txt", "text/plain", []byte("contenu"))
		if attachment.FileName != "测试.txt" {
			t.Errorf("Unicode FileName: got %q, want %q", attachment.FileName, "测试.txt")
		}
		if attachment.Size != int64(len([]byte("contenu"))) {
			t.Errorf("Unicode content size: got %d, want %d", attachment.Size, len([]byte("contenu")))
		}
	})

	t.Run("special characters in mime type", func(t *testing.T) {
		mimeType := "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
		attachment := NewAttachment("att-special", "doc.docx", mimeType, []byte("content"))
		if attachment.MimeType != mimeType {
			t.Errorf("Special MimeType: got %q, want %q", attachment.MimeType, mimeType)
		}
	})

	t.Run("very long ID", func(t *testing.T) {
		longID := "attachment-" + string(make([]byte, 1000))
		attachment := NewAttachment(longID, "test.txt", "text/plain", []byte("content"))
		if attachment.ID != longID {
			t.Error("Long ID should be preserved")
		}
	})
}

func TestAttachment_SizeConsistency(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
	}{
		{"empty", []byte{}},
		{"small", []byte("test")},
		{"medium", make([]byte, 1024)},
		{"large", make([]byte, 1024*1024)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attachment := NewAttachment("att-1", "test.dat", "application/octet-stream", tt.content)
			
			expectedSize := int64(len(tt.content))
			if attachment.Size != expectedSize {
				t.Errorf("Size inconsistency: got %d, want %d", attachment.Size, expectedSize)
			}

			// Verify size is calculated correctly even if content is modified after creation
			originalSize := attachment.Size
			attachment.Content = append(attachment.Content, byte(0))
			
			// Size should still be the original calculated size
			if attachment.Size != originalSize {
				t.Errorf("Size should not change when content is modified externally: got %d, want %d", attachment.Size, originalSize)
			}
		})
	}
}

