package util

import "encoding/base64"

// Base64Encode base64编码
func Base64Encode(input string) string {
	data := []byte(input)
	encodedString := base64.StdEncoding.EncodeToString(data)
	return encodedString
}

// Base64Decode base64解码
func Base64Decode(input string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decodedData), nil
}
