package hash

import (
	"github.com/alexedwards/argon2id"
)

type Argon2Hash struct {
}

func NewArgon2() Hash {
	return Argon2Hash{}
}

func (h Argon2Hash) Hash(plain string) (string, error) {
	return argon2id.CreateHash(plain, argon2id.DefaultParams)
}

func (h Argon2Hash) Verify(hashed, plain string) (bool, error) {
	return argon2id.ComparePasswordAndHash(plain, hashed)
}
