package examples

import (
	"encoding/base64"
)

// Encode base64 encodes a string and returns a string
func Encode(b []byte) string {
	return base64.RawStdEncoding.EncodeToString(b)
}

// Decode base64 decodes a string and returns a string
func Decode(s string) ([]byte, error) {
	out, err := base64.RawStdEncoding.DecodeString(s)
	return out, err
}
