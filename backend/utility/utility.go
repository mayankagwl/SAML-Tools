package utility

import (
	"bytes"
	"encoding/json"
)

var EmptyByte = []byte(`{}`)

func IsValidJson(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, data); err != nil {
		return false
	}
	if bytes.Equal(buffer.Bytes(), EmptyByte) {
		return false
	}
	return true
}
