package sources

type SourceType string

const (
	Bin   SourceType = "bin"
	Card  SourceType = "card"
	Id    SourceType = "id"
	Token SourceType = "token"
)

type (
	SourceRequest interface {
		GetType() SourceType
	}

	MetadataSourceType struct {
		Type SourceType `json:"type,omitempty"`
	}

	requestBinSource struct {
		MetadataSourceType
		Bin string `json:"bin,omitempty"`
	}

	requestCardSource struct {
		MetadataSourceType
		Number string `json:"number,omitempty"`
	}

	requestIdSource struct {
		MetadataSourceType
		Id string `json:"id,omitempty"`
	}

	requestTokenSource struct {
		MetadataSourceType
		Token string `json:"token,omitempty"`
	}
)

func NewRequestBinSource(bin string) *requestBinSource {
	return &requestBinSource{
		Bin:                bin,
		MetadataSourceType: MetadataSourceType{Type: Bin}}
}

func NewRequestCardSource(number string) *requestCardSource {
	return &requestCardSource{
		Number:             number,
		MetadataSourceType: MetadataSourceType{Type: Card}}
}

func NewRequestIdSource(id string) *requestIdSource {
	return &requestIdSource{
		Id:                 id,
		MetadataSourceType: MetadataSourceType{Type: Id}}
}

func NewRequestTokenSource(token string) *requestTokenSource {
	return &requestTokenSource{
		Token:              token,
		MetadataSourceType: MetadataSourceType{Type: Token}}
}

func (c *requestBinSource) GetType() SourceType {
	return c.Type
}

func (c *requestCardSource) GetType() SourceType {
	return c.Type
}

func (c *requestIdSource) GetType() SourceType {
	return c.Type
}

func (c *requestTokenSource) GetType() SourceType {
	return c.Type
}
