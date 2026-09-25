package postgresql_entity

import (
	"time"
)

type AuthenticationAccountGroups struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"column:key;size:50;not null;uniqueIndex"`
	Description string    `gorm:"column:key;size:255"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;index"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type AutherticationAccountGroupsUser struct {
	ID        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"column:userid;not null"`
	GroupId   uint      `gorm:"column:key;size:255"`
	CreatedAt time.Time `gorm:"column:created_at;not null;index"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
