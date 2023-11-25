package database

import (
	"log"

	"github.com/kalleriakronos24/golang-experimental/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(config.AppConfig.DBUrl), &gorm.Config{})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgresSQL")
		return nil
	}
	return db
}
