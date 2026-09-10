package vo

import (
	"errors"
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

func (u Username) Value() string {
	return u.value
}
