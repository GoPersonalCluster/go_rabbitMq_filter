package auth_test

import (
	"encoding/json"
	"net"
	"testing"

	redismodels "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/redis_models"
	"github.com/stretchr/testify/assert"
)

func TestTestNewIPAddress(t *testing.T) {
	ip := net.ParseIP("192.168.1.1")

	authIp := redismodels.NewAuthenticationIp(ip)
	jsonb, err := json.Marshal(authIp)

	if err != nil {
		return
	}

	json.Unmarshal(jsonb, &authIp)

	println(string(jsonb))
	assert.NotNil(t, jsonb)

}
