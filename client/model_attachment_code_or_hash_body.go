package client

import "os"

type AttachmentCodeOrHashBody struct {
	File []*os.File `json:"file,omitempty"`
}
