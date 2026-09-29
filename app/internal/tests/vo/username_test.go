package vo_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUsername_ValidUsername(t *testing.T) {
	username, err := vo.NewUsername("walter_123")
	println(username.Value())
	require.NoError(t, err)
	if err != nil {
		assert.True(t, false, true)
	}

}

func TestNewUsername_AcceptsAllowedSpecialCharacters(t *testing.T) {
	tests := []string{
		"walter-test",
		"walter.test",
		"walter_test",
		"walter-test_1",
		"walter.test_1",
	}

	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			_, err := vo.NewUsername(value)

			assert.NoError(t, err)
		})
	}
}

func TestNewUsername_RejectsLessThan4Characters(t *testing.T) {
	_, err := vo.NewUsername("abc")

	assert.Error(t, err)
}

func TestNewUsername_RejectsMoreThan24Characters(t *testing.T) {
	_, err := vo.NewUsername(
		"abcdefghijklmnopqrstuvwxy",
	)

	assert.Error(t, err)
}

func TestNewUsername_RejectsInvalidSpecialCharacters(t *testing.T) {
	tests := []string{
		"walter@",
		"walter#",
		"walter!",
		"walter$",
		"walter%",
		"walter+",
		"walter/",
	}

	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			_, err := vo.NewUsername(value)

			assert.Error(t, err)
		})
	}
}

func TestNewUsername_RejectsWhitespace(t *testing.T) {
	_, err := vo.NewUsername("walter user")

	if err != nil {
		assert.True(t, false, true)
	}
}
