package vo_test

import (
	"encoding/json"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIpAddressJsonUnmanshal(t *testing.T) {
	tests := []string{
		"127.0.0.1",
		"192.168.0.1",
		"10.0.0.1",
		"8.8.8.8",
	}

	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			ip := net.ParseIP(value)
			jsonb, err := json.Marshal(ip)

			decoded := &net.IP{}
			json.Unmarshal(jsonb, decoded)

			println(string(decoded.String()))

			require.NoError(t, err)
			assert.Equal(t, value, ip)
		})
	}

}
