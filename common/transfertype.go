package common

// TransferType ...
type TransferType string

const (
    // Commission ...
    Commission TransferStatus = "commission"
    // Promotion ...
    Promotion TransferStatus = "promotion"
    // Refund ...
    Refund TransferStatus = "refund"
)

func (c TransferType) String() string {
    return string(c)
}
