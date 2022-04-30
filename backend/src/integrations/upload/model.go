package upload

import (
	"encoding/json"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type (
	// AttachmentKind string representation of attachment types supported
	AttachmentKind string
)

const (
	// AttachmentKindPDF pdf AttachmentKind
	AttachmentKindPDF AttachmentKind = "application/pdf"
	// AttachmentKindImageJPEG image AttachmentKind
	AttachmentKindImageJPEG AttachmentKind = "image/jpeg"
	// AttachmentKindImageJPG image AttachmentKind
	AttachmentKindImageJPG AttachmentKind = "image/jpg"
	// AttachmentKindImagePNG image AttachmentKind
	AttachmentKindImagePNG AttachmentKind = "image/png"
	// AttachmentKindAudio audio AttachmentKind
	AttachmentKindAudio AttachmentKind = "audio"
	// AttachmentKindVideo video AttachmentKind
	AttachmentKindVideo AttachmentKind = "video"
	// ReportFileType report FileType
	ReportFileType FileType = "report"
)

// AttachmentKindMap is handy when we wish to check if a kind exist in a collection
var AttachmentKindMap = map[string]AttachmentKind{
	string(AttachmentKindPDF):       AttachmentKindPDF,
	string(AttachmentKindImageJPEG): AttachmentKindImageJPEG,
	string(AttachmentKindImageJPG):  AttachmentKindImageJPG,
	string(AttachmentKindImagePNG):  AttachmentKindImagePNG,
	string(AttachmentKindAudio):     AttachmentKindAudio,
	string(AttachmentKindVideo):     AttachmentKindVideo,
}

// FileType is a string representation of files types the customer can upload
type FileType string

// FileInput object
type FileInput struct {
	Content []byte
	Name    string
	Kind    string
	URL     string
	Size    int64
	Type    *FileType
}

// FileAttachment object
type FileAttachment struct {
	Kind      string   `json:"kind"`
	URL       string   `json:"url"`
	Size      int64    `json:"size"`
	Extension string   `json:"extension"`
	Name      string   `json:"name,omitempty"`
	Type      FileType `json:"type,omitempty"`
}

// SetAttachments sets Multiple FileAttachments
func SetAttachments(attachments []FileAttachment) (*postgres.Jsonb, error) {
	j, err := json.Marshal(attachments)
	if err != nil {
		return nil, err
	}
	return &postgres.Jsonb{RawMessage: j}, nil
}

// SetAttachment sets Single FileAttachment
func SetAttachment(attachment FileAttachment) (*postgres.Jsonb, error) {
	j, err := json.Marshal(attachment)
	if err != nil {
		return nil, err
	}
	return &postgres.Jsonb{RawMessage: j}, nil
}

// GetAttachments get the attachments in the FileAttachment format
func GetAttachments(attach *postgres.Jsonb) ([]map[string]interface{}, error) {
	var attachments []map[string]interface{}
	b, err := json.Marshal(&attach)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &attachments)
	return attachments, err
}
