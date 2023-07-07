package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(input string) string {
	inputBytes := []byte(input)

	// Create a new SHA256 hash object
	hash := sha256.New()
	// Write the input bytes to the hash object
	hash.Write(inputBytes)

	// Calculate the SHA256 hash
	hashBytes := hash.Sum(nil)

	// Convert the hash bytes to a hexadecimal string
	return hex.EncodeToString(hashBytes)
}
