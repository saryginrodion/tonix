package utils

import (
	"encoding/base64"
)

func ToBase64(s string) (string) {
	encodedString := base64.StdEncoding.EncodeToString([]byte(s))

	return encodedString
}

func FromBase64(base64String string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(base64String)
	return string(res), err
}
