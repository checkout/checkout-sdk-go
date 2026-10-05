package accounts

import (
	"github.com/checkout/checkout-sdk-go/v3/common"
)

// File is a file upload for the Accounts API. SubmitFile (POST /files) sends File and Purpose as a
// multipart request. UploadFile (POST /entities/{entityId}/files) sends only Purpose, as JSON
// (PlatformsFileUpload), and ignores File: the content goes to the upload link in the response.
type File struct {
	// The path to the file to upload (JPEG, PNG or PDF).
	// [Required] for SubmitFile. Not sent by UploadFile, which ignores it; PUT the file content to the
	// upload link UploadFile returns instead.
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
