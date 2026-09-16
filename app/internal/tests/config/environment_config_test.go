package config_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/os_config"
)

func TestGetEnvironmentConfig(t *testing.T) {
	envConfig := os_config.NewEnvironmentConfig()

	if envConfig == nil {
		t.Fatal("expected config, got nil")
	}

}
