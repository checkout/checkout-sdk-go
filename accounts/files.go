package accounts

import (
	"github.com/checkout/checkout-sdk-go/v3/common"
)

// File is a file upload for the Accounts API, sent as a multipart request by SubmitFile
// (POST /files) and UploadFile (POST /entities/{entityId}/files).
type File struct {
	// The path to the file to upload (JPEG, PNG or PDF).
	// [Required]
	File string
	// The purpose of the file upload: the onboarding document the file is for.
	// [Required]
	Purpose common.Purpose
}

func (f *File) GetFile() string {
	return f.File
}

func (f *File) GetPurpose() common.Purpose {
	return f.Purpose
}

func (f *File) GetFieldName() string {
	return "path"
}
