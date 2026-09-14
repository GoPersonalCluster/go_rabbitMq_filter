package redismodels

import (
	"time"
)

type Authentication struct {
	Email     string    `json:"email"`
	Score     float64   `json:"score"`
	UpdatedAt time.Time `json:"updatedat"`
	CreatedAt time.Time `json:"createdat"`
}
