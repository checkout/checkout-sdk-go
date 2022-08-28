package common

// TransferStatus ...
type TransferStatus string

const (
    // TransferPending ...
    TransferPending TransferStatus = "pending"
    // TransferCompleted ...
    TransferCompleted TransferStatus = "completed"
    // TransferRejected ...
    TransferRejected TransferStatus = "rejected"
)

func (c TransferStatus) String() string {
    return string(c)
}
