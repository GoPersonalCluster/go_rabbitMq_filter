package vo_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIPAddress(t *testing.T) {
	tests := []string{
		"127.0.0.1",
		"192.168.0.1",
		"10.0.0.1",
		"8.8.8.8",
	}
	assert.True(t, true)
	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			ip, err := vo.NewIPAddress(value)

			require.NoError(t, err)
			assert.Equal(t, value, ip.String())
		})
	}

}
