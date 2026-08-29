package config

import (
	"fmt"
	"os"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}

	return value
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
	
	JwtSecretKey	 	string
	JwtTokenExpiration 	string 

	RedisAddress	 	string
	RedisPassword	 	string 
	RedisDB 		 	string 
}

func GetHostName() string {
	return "GO_RabbitMqFilter"+NewEnvironmentConfig().HostName
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

		JwtSecretKey:	  	getEnv("jwt_secret_key"),
		JwtTokenExpiration:	getEnv("jwt_token_expiration"),

		RedisAddress:	  	getEnv("redis_addr"),
		RedisPassword:	  	getEnv("redis_password"),
		RedisDB:		  	getEnv("redis_db"),
	}
}
