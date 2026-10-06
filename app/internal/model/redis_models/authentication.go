package redismodels

import (
	"time"
)

type AuthenticationIP struct {
	Score     float64   `json:"score"`
	CreatedAt time.Time `json:"createdat"`
}

func NewAuthenticationIp() *AuthenticationIP {
	return &AuthenticationIP{
		Score:     100,
		CreatedAt: time.Now(),
	}
}

type AuthenticationAccount struct {
	UserId    uint      `json:"userid"`
	CreatedAt time.Time `json:"createdat"`
}

type AuthenticationDevice struct {
	Device string `json:"device"`
	UserId uint   `json:"userid"`
}

/*
-- O Score deve ser superior á 0 para autenticar
-- O score deve possuir um tempo pré definido para ser reiniciado
-- A renovação do token deve ser feita apenas pelo mesmo dispositivo que gerou o token


*/
