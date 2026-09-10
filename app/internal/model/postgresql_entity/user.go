package postgresql_entity

import (
	"time"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/vo"
)

type User struct {
	ID uint `gorm:"primaryKey"`

	Username vo.Username `gorm:"column:username;size:24;not null;uniqueIndex"`
	Email    vo.Email    `gorm:"column:email;size:255;not null;uniqueIndex"`
	Password vo.Password `gorm:"column:password_hash;size:255;not null"`
	Active   bool        `gorm:"not null;default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}

func NewUser(
	username string,
	email string,
	password string,
) (*User, error) {

	usernameVO, err := vo.NewUsername(username)

	if err != nil {
		return nil, err
	}

	emailVO, err := vo.NewEmail(email)

	if err != nil {
		return nil, err
	}

	passwordVO, err := vo.NewPassword(password)

	if err != nil {
		return nil, err
	}

	return &User{
		Username: usernameVO,
		Email:    emailVO,
		Password: passwordVO,
		Active:   true,
	}, nil
}

func (u *User) VerifyPassword(password string) bool {
	return u.Password.Verify(password)
}

func (u *User) ChangePassword(password string) error {
	newPassword, err := vo.NewPassword(password)

	if err != nil {
		return err
	}

	u.Password = newPassword

	return nil
}
