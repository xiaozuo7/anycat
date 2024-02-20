package util

import (
	"fmt"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	input := "anycat"
	encodedString := Base64Encode(input)
	fmt.Println(encodedString)
}

func TestBase64Decode(t *testing.T) {
	input := "YW55Y2F0"
	decodedData, err := Base64Decode(input)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(decodedData)
}
