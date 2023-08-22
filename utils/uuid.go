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

func HashString(s string) string {
	bytes := []byte(s)
	hasher := sha1.New()
	hasher.Write(bytes)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func SoftCompare(a []byte, b []byte) bool {
	hasherA := sha1.New()
	hasherA.Write(a)
	hasherB := sha1.New()
	hasherB.Write(b)
	return hasherA.Size() == hasherB.Size()
}

func HardCompare(a []byte, b []byte) bool {
	hashA := sha1.Sum(a)
	hashB := sha1.Sum(b)
	return hashA == hashB
}
