package util

import (
	"encoding/base64"
	"fmt"
)

func EncodeAnswerToBase64(answer []byte) string {
	return base64.StdEncoding.EncodeToString(answer)
}

func DecodeAnswerFromBase64(encodedAnswer string) ([]byte, error) {
	decodedAnswer, err := base64.StdEncoding.DecodeString(encodedAnswer)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Base64 string: %w", err)
	}
	return decodedAnswer, nil
}
