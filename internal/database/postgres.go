package database

import (
	"log"

	"github.com/shubham-yadavv/go-ecommerce/config"
	"github.com/shubham-yadavv/go-ecommerce/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgres_user = config.GetEnv("POSTGRES_USER")
var postgres_password = config.GetEnv("POSTGRES_PASSWORD")
var postgres_db = config.GetEnv("POSTGRES_DB")

func ConnectPostgres() *gorm.DB {
	dbURL := "postgres://" + postgres_user + ":" + postgres_password + "@localhost:5432/" + postgres_db + "?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	return db
}
