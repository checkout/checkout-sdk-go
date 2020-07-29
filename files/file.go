package files

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"path/filepath"
	"time"

	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/common"
)

type (
	// Request -
	Request struct {
		*QueryParameter
		*FileUpload
	}

	// QueryParameter -
	QueryParameter struct {
	}

	// FileUpload -
	FileUpload struct {
		FileReader  io.Reader
		File        *string
		Purpose     *string
		ContentType *string
	}
)

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// CreateFormFile -
func CreateFormFile(w *multipart.Writer, fieldname string, filename string, contentType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldname, filename))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

// GetBody -
func (f *FileUpload) GetBody() (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if f.Purpose != nil {
		err := writer.WriteField("purpose", StringValue(f.Purpose))
		if err != nil {
			return nil, "", err
		}
	}
	if f.FileReader != nil && f.File != nil && f.ContentType != nil {
		part, err := CreateFormFile(writer, "file", filepath.Base(StringValue(f.File)), StringValue(f.ContentType))
		if err != nil {
			return nil, "", err
		}

		_, err = io.Copy(part, f.FileReader)
		if err != nil {
			return nil, "", err
		}
	}
	err := writer.Close()
	if err != nil {
		return nil, "", err
	}
	return body, writer.Boundary(), nil
}

type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		File           *File                    `json:"file,omitempty"`
	}
	// File -
	File struct {
		ID         string                 `json:"id,omitempty"`
		Filename   string                 `json:"filename,omitempty"`
		Purpose    string                 `json:"purpose,omitempty"`
		Size       uint64                 `json:"size,omitempty"`
		UploadedOn time.Time              `json:"uploaded_on,omitempty"`
		Links      map[string]common.Link `json:"_links"`
	}
)
