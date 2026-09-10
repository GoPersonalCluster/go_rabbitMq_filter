package db

import (
	"fmt"
	"log"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	conf := config.NewEnvironmentConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.PostgresHost,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresPort,
		conf.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
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
	return db
}
