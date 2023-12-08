package base64s

import "encoding/base64"

func EnCode(originalString string) string {
	return base64.StdEncoding.EncodeToString([]byte(originalString))
}

func DeCode(deCodeString string) (string, error) {
	originalString, err := base64.StdEncoding.DecodeString(deCodeString)
	return string(originalString), err
}
