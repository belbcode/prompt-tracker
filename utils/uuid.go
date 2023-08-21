package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

func GenerateRandomID(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func HashName(n string) string {
	bytes := []byte(n)
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
