package common

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"time"

	"github.com/gabriel-vasile/mimetype"

	"github.com/checkout/checkout-sdk-go/errors"
)

type Purpose string

const (
	// Disputes
	DisputesEvidence Purpose = "dispute_evidence"

	// Accounts
	AdditionalDocument           Purpose = "additional_document"
	ArticlesOfAssociation        Purpose = "articles_of_association"
	BankVerification             Purpose = "bank_verification"
	CertifiedAuthorisedSignatory Purpose = "certified_authorised_signatory"
	CompanyOwnership             Purpose = "company_ownership"
	CompanyVerification          Purpose = "company_verification"
	FinancialVerification        Purpose = "financial_verification"
	Identification               Purpose = "identification"
	IdentityVerification         Purpose = "identity_verification"
	TaxVerification              Purpose = "tax_verification"
	ProofOfLegality              Purpose = "proof_of_legality"
	ProofOfPrincipalAddress      Purpose = "proof_of_principal_address"
	ShareholderStructure         Purpose = "shareholder_structure"
	ProofOfResidentialAddress    Purpose = "proof_of_residential_address"
	ProofOfRegistration          Purpose = "proof_of_registration"
)

type (
	FileUpload interface {
		GetFile() string
		GetPurpose() Purpose
		GetFieldName() string
	}

	File struct {
		File    string
		Purpose Purpose
	}

	FileUploadRequest struct {
		W *multipart.Writer
		B *bytes.Buffer
	}

	FileResponse struct {
		HttpMetadata HttpMetadata
		Id           string          `json:"id,omitempty"`
		Filename     string          `json:"filename,omitempty"`
		Purpose      Purpose         `json:"purpose,omitempty"`
		Size         uint64          `json:"size,omitempty"`
		UploadedOn   *time.Time      `json:"uploaded_on,omitempty"`
		Links        map[string]Link `json:"_links"`
	}
)

func BuildFileUploadRequest(upload FileUpload) (*FileUploadRequest, error) {
	if err := validateFile(upload); err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	file, err := os.Open(upload.GetFile())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	contentType, err := mimetype.DetectFile(upload.GetFile())
	if err != nil {
		return nil, err
	}

	part, err := createFormFile(writer, upload.GetFieldName(), filepath.Base(file.Name()), contentType.String())
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("purpose", string(upload.GetPurpose()))
	if err != nil {
		return nil, err
	}

	return &FileUploadRequest{
		W: writer,
		B: body,
	}, nil
}

func validateFile(f FileUpload) error {
	if f.GetFile() == "" {
		return errors.BadRequestError("Invalid file name")
	}
	if f.GetPurpose() == "" {
		return errors.BadRequestError("Invalid purpose")
	}

	return nil
}

func createFormFile(w *multipart.Writer, fieldName string, fileName string, contentType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		EscapeQuotes(fieldName),
		EscapeQuotes(fileName)))
	h.Set("Content-Type", EscapeQuotes(contentType))

	return w.CreatePart(h)
}

func (f *File) GetFile() string {
	return f.File
}

func (f *File) GetPurpose() Purpose {
	return f.Purpose
}

func (f *File) GetFieldName() string {
	return "file"
}
