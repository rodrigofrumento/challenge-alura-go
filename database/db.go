package database

import (
	"log"

	"github.com/rodrigofrumento/challenge-alura-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConn() {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=videos sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Panic("Connection to DB error.")
	}
	DB.AutoMigrate(&models.Video{})
}
