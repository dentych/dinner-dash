package security

import (
	"crypto/rand"
	"strings"
)

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateSession() (string, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	mod := byte(len(letters))

	var sb strings.Builder
	for _, value := range b {
		index := value % mod
		sb.WriteByte(letters[index])
	}

	return sb.String(), nil
}
