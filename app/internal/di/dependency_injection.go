package di

import (
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DbConnectionDi struct {
	Con *gorm.DB
	ID  uuid.UUID
}

func NewDbConnectionDi() *DbConnectionDi {
	return &DbConnectionDi{
		Con: db.GetDbConnection(),
		ID:  uuid.New(),
	}
}

type dependency_injection struct {
	DbConnection *DbConnectionDi
}

func (di *dependency_injection) GetDbConnection() *DbConnectionDi {
	if di != nil {
		return di.DbConnection
	}
	return NewDbConnectionDi()

}
