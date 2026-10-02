package redismodels

import (
	"time"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
)

type AuthenticationIP struct {
	IP    vo.IPAddress `json:"ip"`
	Score float64      `json:"score"`
}

func NewAuthenticationIp(ip vo.IPAddress) *AuthenticationIP {
	return &AuthenticationIP{
		IP:    ip,
		Score: 100,
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
