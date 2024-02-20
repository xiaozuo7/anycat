package service

import "anycat/util"

func Base64Encode(content string) string {
	res := util.Base64Encode(content)
	return res
}

func Base64Decode(content string) (string, error) {
	res, err := util.Base64Decode(content)
	if err != nil {
		return "", err
	}
	return res, nil
}
