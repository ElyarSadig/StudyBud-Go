package main

import (
	"log"

	"github.com/elyarsadig/studybud-go/migrations"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger, err := logger.New(logger.JSON, logger.DebugLevel)
	if err != nil {
		log.Fatal(err)
	}
	logger.Debug("This is a start!", "start", "end")
	db := initDB()
	err = migrations.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
}

func initDB() (*gorm.DB) {
	config := postgres.New(postgres.Config{
		DSN: "user=admin password=password dbname=studybud_db port=5432 sslmode=disable",
	})
	db, err := gorm.Open(config, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
