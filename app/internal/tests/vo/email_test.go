package vo_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail_ValidEmails(t *testing.T) {
	tests := []string{
		"walter@example.com",
		"walter.matsuda@example.com",
		"walter+api@example.com",
		"walter123@example.co.uk",
	}

	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			t.Parallel()

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
			t.Parallel()

			_, err := vo.NewEmail(value)
			assert.Error(t, err)
		})
	}
}

func TestNewEmail_RejectsDisplayName(t *testing.T) {
	_, err := vo.NewEmail(
		"Walter <walter@example.com>",
	)

	assert.Error(t, err)
}
