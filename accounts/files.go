package accounts

import "github.com/checkout/checkout-sdk-go-beta/common"

type File struct {
	File    string
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
