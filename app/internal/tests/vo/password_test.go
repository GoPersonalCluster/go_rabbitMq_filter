package vo_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPassword_VerifyValidPassword(t *testing.T) {
	password := "my-secret-password"

	passwordVO, err := vo.NewPassword(password)

	require.NoError(t, err)

	assert.True(
		t,
		passwordVO.Verify(password),
	)
}
