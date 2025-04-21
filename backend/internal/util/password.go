package util

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"strings"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var (
	p *params = &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
)

func GenerateFromPassword(password string) (string, string, error) {
	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", "", err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Encode for DB/Comparison usage
	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)

	return encodedHash, encodedSalt, nil
}

func GenerateFromPasswordWithSalt(password string, saltString string) (string, error) {
	salt, err := base64.StdEncoding.DecodeString(saltString)
	if err != nil {
		return "", err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Encode for DB/Comparison usage
	encodedHash := base64.StdEncoding.EncodeToString(hash)

	return encodedHash, nil
}

func SplitPasswordSalt(password string) (string, string) {
	var split []string = strings.Split(password, ":")
	// idx 0 == salt
	// idx 1 == hashed password
	return split[0], split[1]
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
