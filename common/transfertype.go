package common

// TransferType ...
type TransferType string

const (
    // Commission ...
    Commission TransferType = "commission"
    // Promotion ...
    Promotion TransferType = "promotion"
    // Refund ...
    Refund TransferType = "refund"
)

func (c TransferType) String() string {
    return string(c)
}
