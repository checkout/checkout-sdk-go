package common

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
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

func Unmarshal(metadata *HttpMetadata, responseMapping interface{}) error {
	if len(metadata.ResponseBody) == 0 {
		addHttpMetadata(metadata, responseMapping)
		return nil
	}

	if isContentTypeText(metadata) {
		addContent(metadata.ResponseBody, responseMapping)
	} else {
		if err := json.Unmarshal(metadata.ResponseBody, &responseMapping); err != nil {
			return err
		}
	}

	addHttpMetadata(metadata, responseMapping)

	return nil
}

func addHttpMetadata(metadata *HttpMetadata, response interface{}) {
	v := reflect.ValueOf(response).Elem().FieldByName("HttpMetadata")
	if v.IsValid() {
		v.Set(reflect.ValueOf(*metadata))
	}
}

func addContent(content []byte, response interface{}) {
	v := reflect.ValueOf(response).Elem().FieldByName("Content")
	if v.IsValid() {
		v.Set(reflect.ValueOf(string(content)))
	}
}

func isContentTypeText(metadata *HttpMetadata) bool {
	return strings.HasPrefix(metadata.Headers.Header.Get("Content-Type"), "text/csv")
}
