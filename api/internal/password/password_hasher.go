package password

import (
	"bytes"
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

type PasswordHasher interface {
	HashPass(plainPass string) []byte
	CheckPass(passHash []byte, plain string) bool
}

type PasswordHasherImpl struct {
}

func NewPasswordHasher() *PasswordHasherImpl {
	return &PasswordHasherImpl{}
}

func (ph *PasswordHasherImpl) HashPass(plainPass string) []byte {
	salt := make([]byte, 8)
	rand.Read(salt)
	hash := ph.hashPassInternal(salt, plainPass)
	return hash
}

func (ph *PasswordHasherImpl) CheckPass(passHash []byte, plain string) bool {
	salt := passHash[:8]
	newSalt := make([]byte, len(salt))
	copy(newSalt, salt)
	inputPassHash := ph.hashPassInternal(newSalt, plain)

	return bytes.Equal(passHash, inputPassHash)
}

func (ph *PasswordHasherImpl) hashPassInternal(salt []byte, plainPass string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPass), salt, 1, 64*1024, 4, 32)
	return hashedPass
}
