package vo_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmail_ValidEmail(t *testing.T) {
	email, err := vo.NewEmail(
		"walter@example.com",
	)

	require.NoError(t, err)
	assert.Equal(t, "walter@example.com", email.Value())
}

func TestNewEmail_ValidEmails(t *testing.T) {
	tests := []string{
		"walter@example.com",
		"walter.matsuda@example.com",
		"walter+api@example.com",
		"walter123@example.co.uk",
	}

	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			_, err := vo.NewEmail(value)

			assert.NoError(t, err)
		})
	}
}

func TestNewEmail_RejectsInvalidEmails(t *testing.T) {
	tests := []string{
		"",
		"walter",
		"walter@",
		"@example.com",
		"walter@example",
		"walter @example.com",
		"walter@example.com extra",
	}

	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			_, err := vo.NewEmail(value)
			assert.True(t, err != nil)

		})
	}
}

func TestNewEmail_RejectsDisplayName(t *testing.T) {
	_, err := vo.NewEmail(
		"Walter <walter@example.com>",
	)

	assert.Error(t, err)
}
