package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID generates a random 32-character hex string unique ID
func GenerateID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
