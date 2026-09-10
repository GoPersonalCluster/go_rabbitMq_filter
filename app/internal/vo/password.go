package vo

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash string
}

func NewPassword(plainText string) (Password, error) {
	if len(plainText) < 8 {
		return Password{}, errors.New("password must contain at least 8 characters")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(plainText),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return Password{}, err
	}

	return Password{
		hash: string(hash),
	}, nil
}

func NewPasswordFromHash(hash string) (Password, error) {
	if hash == "" {
		return Password{}, errors.New("password hash cannot be empty")
	}

	return Password{
		hash: hash,
	}, nil
}

func (p Password) Verify(plainText string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(p.hash),
		[]byte(plainText),
	) == nil
}

func (p Password) Hash() string {
	return p.hash
}
