package files

import (
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
		File    string `json:"file,omitempty"`
		Purpose string `json:"purpose,omitempty"`
	}
)
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
