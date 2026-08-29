package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/repository"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	redis *repository.RedisRepository
}

func NewUserService(
	redis *repository.RedisRepository,
) *UserService {

	return &UserService{
		redis: redis,
	}
}

func (s *UserService) GetUser(
	ctx context.Context,
	id int,
) (*User, error) {

	key := "user:" + string(rune(id))

	// Primeiro tenta buscar no Redis.
	value, err := s.redis.Get(ctx, key)

	if err == nil {
		var user User

		if err := json.Unmarshal(
			[]byte(value),
			&user,
		); err == nil {

			return &user, nil
		}
	}

	// Simulando acesso ao banco.
	user := &User{
		ID:    id,
		Name:  "Walter",
		Email: "walter@example.com",
	}

	// Serializa para JSON.
	data, err := json.Marshal(user)

	if err != nil {
		return nil, err
	}

	// Cache por 5 minutos.
	err = s.redis.Set(
		ctx,
		key,
		string(data),
		5*time.Minute,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}