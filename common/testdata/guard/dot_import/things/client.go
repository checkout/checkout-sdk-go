package things

import . "github.com/checkout/checkout-sdk-go/v3/common"

func Path(id string) string {
	return BuildPath("things", id)
}
