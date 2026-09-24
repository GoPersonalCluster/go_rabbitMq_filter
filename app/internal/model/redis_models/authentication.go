package redismodels

import (
	"time"
)

type AuthenticationIP struct {
	IP        string    `json:"email"`
	Score     float64   `json:"score"`
	UpdatedAt time.Time `json:"updatedat"`
	CreatedAt time.Time `json:"createdat"`
}
type AuthenticationAccountGroups struct {
}
type AuthenticationIpGroup struct {
}
type AuthenticationAccount struct {
}
type AuthenticationDevice struct {
}
type AuthenticationNetworkInterface struct {
}
