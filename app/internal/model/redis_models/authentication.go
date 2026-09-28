package redismodels

import (
	"time"
)

type AuthenticationIP struct {
	IP        string    `json:"ip"`
	Score     float64   `json:"score"`
	UpdatedAt time.Time `json:"updatedat"`
	CreatedAt time.Time `json:"createdat"`
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
