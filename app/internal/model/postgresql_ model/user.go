package postgresql_model

import (
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"size:100;not null;uniqueIndex"`
	Email        string `gorm:"size:255;not null;uniqueIndex"`
	PasswordHash string `gorm:"column:password_hash;size:255;not null"`
	Active       bool   `gorm:"not null;default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	username string,
	email string,
	password string,
) (*User, error) {

	passwordVO, err := vo.NewPassword(password)

	if err != nil {
		return nil, err
	}

	return &User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordVO.Hash(),
		Active:       true,
	}, nil
}

func (u *User) VerifyPassword(password string) bool {
	passwordVO, err := vo.NewPasswordFromHash(
		u.PasswordHash,
	)

	if err != nil {
		return false
	}

	return passwordVO.Verify(password)
}
