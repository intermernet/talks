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

// Unordered returns the input slice in a random order
func Unordered(i []int) []int {
	m := make(map[int]int)
	out := make([]int, 0)
	for a, b := range i {
		m[a] = b
	}
	for _, v := range m {
		out = append(out, v)
	}
	return out
}
