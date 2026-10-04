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

type AuthenticationNetworkInterface struct {
	HttpHeader string `json:"httpheader"`
}
