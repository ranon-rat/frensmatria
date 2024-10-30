package common

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(input string) string {
	// Crear un hash SHA-256
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
