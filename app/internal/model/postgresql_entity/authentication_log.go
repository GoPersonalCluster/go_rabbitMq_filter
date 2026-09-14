package postgresql_entity

import (
	"time"
)

type AuthenticationLog struct {
	ID                  uint                  `gorm:"primaryKey"`
	Severity            uint                  `gorm:"column:severity;not null"`
	UserId              uint                  `gorm:"column:severity;not null"`
	User                User                  `gorm:"foreignKey:UserId;references:Key"`
	AuthenticationLogId uint                  `gorm:"column:AuthenticationLogId;not null;index"`
	Code                AuthenticationLogCode `gorm:"foreignKey:AuthenticationLogId;references:Id"`
	CreatedAt           time.Time             `gorm:"column:created_at;not null;index"`
}

func (AuthenticationLog) TableName() string {
	return "authentication_logs"
}

type AuthenticationLogCode struct {
	ID          uint      `gorm:"primaryKey"`
	Key         string    `gorm:"column:key;size:50;not null;uniqueIndex"`
	Description string    `gorm:"column:description;size:255;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
}

func (AuthenticationLogCode) TableName() string {
	return "authentication_log_codes"
}
