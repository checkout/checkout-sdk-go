package things

import _ "unsafe"

//go:linkname buildPath github.com/checkout/checkout-sdk-go/v3/common.BuildPath
func buildPath(segments ...string) string
