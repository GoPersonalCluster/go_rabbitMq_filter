package os_config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}

	return value
}
func getEnvFloat(key string) float64 {
	value, exists := os.LookupEnv(key)

	if !exists || strings.TrimSpace(value) == "" {
		return 0
	}

	floatVal, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		panic(fmt.Sprintf("environment variable %s is invalid", key))
	}

	return floatVal
}

func getEnvFloat64(key string) float64 {
	return getEnvFloat(key)
}

type EnvironmentConfig struct {
	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     string

	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	HostName         string

	JwtSecretKey       string
	JwtTokenExpiration string

	RedisAddress                     string
	RedisPassword                    string
	RedisDB                          string
	AuthenticationScoreFailedAttempt float64
	AuthenticationSpamScore          float64
}

func GetHostName() string {
	return "GO_RabbitMqFilter" + NewEnvironmentConfig().HostName
}

func NewEnvironmentConfig() *EnvironmentConfig {
	return &EnvironmentConfig{
		RabbitMQUsername: getEnv("rabbitmq_username"),
		RabbitMQPassword: getEnv("rabbitmq_password"),
		RabbitMQHost:     getEnv("rabbitmq_host"),
		RabbitMQPort:     getEnv("rabbitmq_port"),

		PostgresHost:     getEnv("postgres_host"),
		PostgresPort:     getEnv("postgres_port"),
		PostgresDB:       getEnv("postgres_db"),
		PostgresUser:     getEnv("postgres_user"),
		PostgresPassword: getEnv("postgres_password"),
		HostName:         getEnv("host_name"),

		JwtSecretKey:       getEnv("jwt_secret_key"),
		JwtTokenExpiration: getEnv("jwt_token_expiration"),

		RedisAddress:                     getEnv("redis_addr"),
		RedisPassword:                    getEnv("redis_password"),
		RedisDB:                          getEnv("redis_db"),
		AuthenticationScoreFailedAttempt: getEnvFloat("authentication_score_failed_attempt"),
		AuthenticationSpamScore:          getEnvFloat("authentication_spam_score"),
	}
}
