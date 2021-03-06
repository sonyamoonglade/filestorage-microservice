package headers

import (
	"encoding/json"
	"github.com/fatih/structs"
)

const (
	XFileExt       = "x-file-ext"
	XFileName      = "x-file-name"
	XHmacSignature = "x-hmac-signature"
	XDestination   = "x-destination"
	XHeaders       = "x-headers"
)

type Headers map[string]string

//todo: diff headers for diff requests
type PutHeaders struct {
	XFileExt       string `header:"x-file-ext"`
	XFileName      string `header:"x-file-name"`
	XHmacSignature string `header:"x-hmac-signature"`
	XDestination   string `header:"x-destination"`
}

func Validate(v interface{}) bool {
	result := structs.HasZero(v)
	return !result
}
func Decode(v interface{}, out *map[string]string) error {
	b, _ := json.Marshal(v)
	err := json.Unmarshal(b, out)
	if err != nil {
		return err
	}
	return nil
}
