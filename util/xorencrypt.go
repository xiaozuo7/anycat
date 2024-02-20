package util

import "encoding/base64"

const key = 18

// XorEncryptStr 异或加密
func XorEncryptStr(str string) string {
	encrypt := base64.StdEncoding.EncodeToString([]byte(str))
	var tmp byte
	var text string
	for i := 0; i < len(encrypt); i++ {
		tmp = encrypt[i] ^ key
		if isLegalFlag(tmp) {
			text += string(tmp)
		} else {
			text += string(encrypt[i])
		}
	}
	return text
}

// XorDecryptStr 异或解密
func XorDecryptStr(str string) (string, error) {
	var tmp byte
	var text string
	for i := 0; i < len(str); i++ {
		tmp = str[i] ^ key
		if isLegalFlag(tmp) {
			text += string(tmp)
		} else {
			text += string(str[i])
		}
	}
	decrypt, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	return string(decrypt), nil
}

func isLegalFlag(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}
	if b >= 65 && b <= 90 {
		return true
	}
	if b >= 97 && b <= 122 {
		return true
	}
	return false
}
