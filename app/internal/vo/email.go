package vo

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

type Email struct {
	value string
}

func (e Email) Value() (driver.Value, error) {
	return e.value, nil
}

func (e *Email) Scan(value any) error {
	var email string

	switch v := value.(type) {
	case string:
		email = v

	case []byte:
		email = string(v)

	default:
		return fmt.Errorf("cannot scan %T into Email", value)
	}

	emailVO, err := NewEmail(email)
	if err != nil {
		return err
	}

	*e = emailVO

	return nil
}
func NewEmail(value string) (Email, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return Email{}, errors.New("email cannot be empty")
	}

	err := validateEmail(value)

	if err != nil {
		return Email{}, errors.New("invalid email")
	}

	return Email{
		value: value,
	}, nil
}

func validateEmail(value string) error {
	address, err := mail.ParseAddress(value)
	if err != nil {
		return fmt.Errorf("invalid email")
	}

	if address.Address != value {
		return fmt.Errorf("invalid email")
	}

	parts := strings.Split(address.Address, "@")
	if len(parts) != 2 {
		return fmt.Errorf("invalid email")
	}

	domain := parts[1]

	if !strings.Contains(domain, ".") {
		return fmt.Errorf("invalid email domain")
	}

	return nil
}
