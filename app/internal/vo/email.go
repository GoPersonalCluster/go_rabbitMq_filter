package vo

import (
	"errors"
	"net/mail"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return Email{}, errors.New("email cannot be empty")
	}

	address, err := mail.ParseAddress(value)

	if err != nil {
		return Email{}, errors.New("invalid email")
	}

	if address.Address != value {
		return Email{}, errors.New("invalid email")
	}

	return Email{
		value: value,
	}, nil
}

func (e Email) String() string {
	return e.value
}

func (e Email) Value() string {
	return e.value
}
