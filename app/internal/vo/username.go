package vo

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"unicode/utf8"
)

type Username struct {
	value string
}

func NewUsername(value string) (Username, error) {
	length := utf8.RuneCountInString(value)

	if length < 4 || length > 24 {
		return Username{}, errors.New(
			"username must contain between 4 and 24 characters",
		)
	}

	for _, char := range value {
		if isAllowedUsernameCharacter(char) {
			continue
		}

		return Username{}, errors.New(
			"username contains invalid characters",
		)
	}

	return Username{
		value: value,
	}, nil
}
func (u Username) Value() (driver.Value, error) {
	return u.value, nil
}

func (u *Username) Scan(value any) error {
	switch v := value.(type) {
	case string:
		username, err := NewUsername(v)
		if err != nil {
			return err
		}

		*u = username
		return nil

	case []byte:
		username, err := NewUsername(string(v))
		if err != nil {
			return err
		}

		*u = username
		return nil

	default:
		return fmt.Errorf("cannot scan %T into Username", value)
	}
}
func isAllowedUsernameCharacter(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		char == '-' ||
		char == '.' ||
		char == '_'
}

func (u Username) String() string {
	return u.value
}
