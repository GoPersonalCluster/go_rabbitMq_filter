package postgresql_entity

import (
	"time"
)

type AuthenticationAccountGroups struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"column:name;size:50;not null;uniqueIndex"`
	Description string    `gorm:"column:description;size:255"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;index"`
	Enabled     bool      `gorm:"column:enabled;not null;default:true"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null"`
}

func (AuthenticationAccountGroups) TableName() string {
	return "authentication_account_groups"
}

type AutherticationAccountGroupsUser struct {
	ID        uint                        `gorm:"primaryKey"`
	UserId    uint                        `gorm:"column:userid;not null;uniqueIndex:idx_user_group"`
	GroupId   uint                        `gorm:"column:groupid;not null;uniqueIndex:idx_user_group"`
	Group     AuthenticationAccountGroups `gorm:"foreignKey:GroupId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Enabled   bool                        `gorm:"column:enabled;not null;default:true"`
	CreatedAt time.Time                   `gorm:"column:created_at;not null;index"`
	UpdatedAt time.Time                   `gorm:"column:updated_at;not null"`
}

func (AutherticationAccountGroupsUser) TableName() string {
	return "authentication_account_groups_user"
}
