package common

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/checkout/checkout-sdk-go/v3/errors"
	"github.com/checkout/checkout-sdk-go/v3/internal/jsondepth"
)

type (
	TypeMapping struct {
		Type        string `json:"type"`
		ControlType string `json:"control_type"`
	}
)

func Marshal(request interface{}) (*bytes.Buffer, error) {
	if request != nil {
		marshal, err := json.Marshal(request)
		return bytes.NewBuffer(marshal), err
	}
	return new(bytes.Buffer), nil
}

// Unmarshal decodes a successful response into responseMapping, which must be a non-nil pointer.
// A text/csv or application/pdf body goes to the Content field of the response (a string, a
// []byte or an interface); any other body is decoded as JSON, and a JSON null leaves the zero
// value. The HttpMetadata field, when the response has one, is then set from metadata.
//
// When the body cannot be decoded it returns errors.CheckoutDecodingError, which carries the status
// code and the media type only. JSON nested deeper than jsondepth.MaxDepth is not decoded.
func Unmarshal(metadata *HttpMetadata, responseMapping interface{}) error {
	if metadata == nil {
		return errors.CheckoutDecodingError{}
	}
	mediaType := responseMediaType(metadata)
	fail := func() error {
		return errors.CheckoutDecodingError{StatusCode: metadata.StatusCode, ContentType: mediaType}
	}

	target := reflect.ValueOf(responseMapping)
	if !target.IsValid() || target.Kind() != reflect.Ptr || target.IsNil() {
		return fail()
	}

	body := metadata.ResponseBody
	if len(body) != 0 {
		switch mediaType {
		case "text/csv":
			if !setContent(target, body, reflect.ValueOf(string(body))) {
				return fail()
			}
		case "application/pdf":
			if !setContent(target, body, reflect.ValueOf(append([]byte(nil), body...))) {
				return fail()
			}
		default:
			if !decodeJSON(body, responseMapping) {
				return fail()
			}
		}
	}

	if !setHttpMetadata(target, metadata) {
		return fail()
	}
	return nil
}

// decodeJSON decodes body into responseMapping and reports whether it succeeded. A JSON null leaves
// responseMapping as it is. It reports false when the body is nested deeper than jsondepth.MaxDepth,
// when decoding returns an error, and when a custom UnmarshalJSON does not return normally.
func decodeJSON(body []byte, responseMapping interface{}) (ok bool) {
	if string(bytes.TrimSpace(body)) == "null" {
		return true
	}
	if !jsondepth.WithinLimit(body) {
		return false
	}
	returned := false
	defer func() {
		if !returned {
			_ = recover()
			ok = false
		}
	}()
	err := json.Unmarshal(body, responseMapping)
	returned = true
	return err == nil
}

// responseMediaType returns the lower-case media type of the response, without parameters, or ""
// when there is none. A value that is not a plain media type token is returned as "invalid", so the
// result can read "invalid" for a header the server sent in another form.
func responseMediaType(metadata *HttpMetadata) string {
	if metadata.Headers == nil || metadata.Headers.Header == nil {
		return ""
	}
	value := metadata.Headers.Header.Get("Content-Type")
	if i := strings.IndexByte(value, ';'); i >= 0 {
		value = value[:i]
	}
	value = strings.ToLower(strings.TrimSpace(value))
	if len(value) > 127 {
		return "invalid"
	}
	for i := 0; i < len(value); i++ {
		if !isMediaTypeByte(value[i]) {
			return "invalid"
		}
	}
	return value
}

// isMediaTypeByte reports whether c can appear in a lower-case media type token.
func isMediaTypeByte(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || strings.IndexByte("/!#$&^_.+-", c) >= 0
}

// setContent sets the Content field of the struct target points to: a string field gets the body
// as text, a []byte field a copy of it, an empty interface field the given value. It reports false
// when there is no such field or it has another type.
func setContent(target reflect.Value, body []byte, value reflect.Value) bool {
	field, ok := settableField(target, "Content")
	if !ok {
		return false
	}
	switch {
	case field.Kind() == reflect.String:
		field.SetString(string(body))
	case field.Kind() == reflect.Slice && field.Type().Elem().Kind() == reflect.Uint8:
		field.SetBytes(append([]byte(nil), body...))
	case field.Kind() == reflect.Interface && field.NumMethod() == 0:
		field.Set(value)
	default:
		return false
	}
	return true
}

// httpMetadataType is the type the HttpMetadata field of a response must have.
var httpMetadataType = reflect.TypeOf(HttpMetadata{})

// setHttpMetadata sets the HttpMetadata field of the struct target points to. A target that is not
// a struct, or has no settable HttpMetadata field, is left as it is; a field of another type
// reports false.
func setHttpMetadata(target reflect.Value, metadata *HttpMetadata) bool {
	field, ok := settableField(target, "HttpMetadata")
	if !ok {
		return true
	}
	if field.Type() != httpMetadataType {
		return false
	}
	field.Set(reflect.ValueOf(*metadata))
	return true
}

// settableField follows target through its pointers to a struct and returns its exported field
// with that name. It reports false when a pointer is nil, the value is not a struct, there is no
// such field, or the field sits behind a nil embedded pointer or cannot be set.
func settableField(target reflect.Value, name string) (reflect.Value, bool) {
	v := target
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflect.Value{}, false
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}
	structField, ok := v.Type().FieldByName(name)
	if !ok {
		return reflect.Value{}, false
	}
	for i, index := range structField.Index {
		if i > 0 {
			if v.Kind() == reflect.Ptr {
				if v.IsNil() {
					return reflect.Value{}, false
				}
				v = v.Elem()
			}
		}
		v = v.Field(index)
	}
	return v, v.CanSet()
}
