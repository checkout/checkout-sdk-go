package accounts

import (
	"github.com/checkout/checkout-sdk-go/v3/common"
)

// File is a file upload for the Accounts API.
//   - SubmitFile (POST /files) sends File and Purpose as a multipart request; File is the path to
//     the file to upload (JPEG, PNG or PDF).
//   - UploadFile (POST /entities/{entityId}/files) sends only Purpose, as the JSON body the endpoint
//     defines, and returns an upload link for the file content; File is ignored there.
type File struct {
	// The path to the file to upload. SubmitFile only.
	// [Required] for SubmitFile
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
