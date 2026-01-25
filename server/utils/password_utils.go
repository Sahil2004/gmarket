package utils

import (
	"encoding/base64"

	"github.com/jaevor/go-nanoid"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, string) {
	saltGen, _ := nanoid.Standard(8)
	salt := saltGen()
	hashBytes := argon2.Key([]byte(password), []byte(salt), 3, 32*1024, 4, 32)
	passwordHash := base64.StdEncoding.EncodeToString(hashBytes)
	return passwordHash, salt
}

func ValidatePassword(password, salt, passwordHash string) bool {
	hashBytes := argon2.Key([]byte(password), []byte(salt), 3, 32*1024, 4, 32)
	computedHash := base64.StdEncoding.EncodeToString(hashBytes)
	return computedHash == passwordHash
}