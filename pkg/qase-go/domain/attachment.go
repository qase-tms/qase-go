package domain

import (
	"fmt"
	"strings"
)

// Attachment represents a file attachment matching the JavaScript interface
type Attachment struct {
	ID       string  `json:"id"`
	FileName string  `json:"file_name"`
	MimeType string  `json:"mime_type"`
	FilePath *string `json:"file_path"` // nullable pointer
	Content  []byte  `json:"content"`   // equivalent to string | Buffer in JS
	Size     int64   `json:"size"`
}

// NewAttachment creates a new attachment instance
func NewAttachment(id, fileName, mimeType string, content []byte) *Attachment {
	return &Attachment{
		ID:       id,
		FileName: fileName,
		MimeType: mimeType,
		Content:  content,
		Size:     int64(len(content)),
		FilePath: nil,
	}
}

// SetFilePath sets the file path
func (a *Attachment) SetFilePath(path string) {
	a.FilePath = &path
}

// GetFilePath returns the file path if set
func (a *Attachment) GetFilePath() string {
	if a.FilePath != nil {
		return *a.FilePath
	}
	return ""
}

// HasFilePath returns true if file path is set
func (a *Attachment) HasFilePath() bool {
	return a.FilePath != nil
}

// String implements the Stringer interface for Attachment
func (a *Attachment) String() string {
	var parts []string

	if a.ID != "" {
		parts = append(parts, fmt.Sprintf("ID: %q", a.ID))
	}
	if a.FileName != "" {
		parts = append(parts, fmt.Sprintf("FileName: %q", a.FileName))
	}
	if a.MimeType != "" {
		parts = append(parts, fmt.Sprintf("MimeType: %q", a.MimeType))
	}
	if a.FilePath != nil {
		parts = append(parts, fmt.Sprintf("FilePath: %q", *a.FilePath))
	}
	if len(a.Content) > 0 {
		parts = append(parts, fmt.Sprintf("Content: %d bytes", len(a.Content)))
	}
	parts = append(parts, fmt.Sprintf("Size: %d", a.Size))

	return fmt.Sprintf("Attachment{%s}", strings.Join(parts, ", "))
}
