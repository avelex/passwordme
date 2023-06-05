package aes

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptDecrypt(t *testing.T) {
	r := require.New(t)

	text := "Hello, world!"
	key := "test"
	hash := sha256.Sum256([]byte(key))

	encrypted, err := Encrypt([]byte(text), hash[:])
	r.NoError(err)

	decrypted, err := Decrypt(encrypted, hash[:])
	r.NoError(err)

	r.Equal(text, string(decrypted))
}
