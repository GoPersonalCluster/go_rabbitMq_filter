package auth_test

import (
	"encoding/json"
	"testing"

	redismodels "github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/redis_models"
	"github.com/stretchr/testify/assert"
)

func TestTestNewIPAddress(t *testing.T) {

	authIp := redismodels.NewAuthenticationIp()
	jsonb, err := json.Marshal(authIp)

	if err != nil {
		return
	}

	json.Unmarshal(jsonb, &authIp)

	println(string(jsonb))
	assert.NotNil(t, jsonb)

}
