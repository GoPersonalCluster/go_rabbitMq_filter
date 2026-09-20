package vo

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func NewPassword(plainText string) (Password, error) {
	if len(plainText) < 8 {
		return Password{}, errors.New("password must contain at least 8 characters")
	}

	value, err := bcrypt.GenerateFromPassword(
		[]byte(plainText),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return Password{}, err
	}

	return Password{
		value: string(value),
	}, nil
}
func (p Password) Value() (driver.Value, error) {
	return p.value, nil
}
func NewPasswordFromvalue(value string) (Password, error) {
	if value == "" {
		return Password{}, errors.New("password value cannot be empty")
	}

	return Password{
		value: value,
	}, nil
}

func (p Password) Verify(plainText string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(p.value),
		[]byte(plainText),
	) == nil
}

func (p *Password) Scan(value any) error {
	var password string

	switch v := value.(type) {
	case string:
		password = v

	case []byte:
		password = string(v)

	default:
		return fmt.Errorf("cannot scan %T into Password", value)
	}

	passwordVO, err := NewPasswordFromvalue(password)
	if err != nil {
		return err
	}

	*p = passwordVO

	return nil
}
