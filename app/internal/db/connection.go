package db

import (
	"fmt"
	"log"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/model/postgresql_entity"
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/os_config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	conf := os_config.NewEnvironmentConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.PostgresHost,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDB,
		conf.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", dsn, err)
	}

	fmt.Println("Database connected successfully")

	// Test the underlying database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
		panic("database configuration is invalid")
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("database is not reachable:", err)
		panic("database configuration is unreachable")
	}

	err = db.AutoMigrate(
		&postgresql_entity.AuthenticationLogCode{},
		&postgresql_entity.AuthenticationLog{},
		&postgresql_entity.User{},
	)

	if err != nil {
		log.Fatal("database is not reachable:", err)
		panic("error during database migration")
	}
	return db
}
