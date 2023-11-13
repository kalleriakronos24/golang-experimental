package services

import (
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	"github.com/kalleriakronos24/mygoapp2nd/config"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
	master "github.com/kalleriakronos24/mygoapp2nd/models/master"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Handler HandlerFunc

type HandlerFunc interface {
	AuthenticateUser(credentials dto.UserLogin) (token string, err error)

	// User Handlers
	RegisterUser(credentials dto.UserSignup) (err error)
	RetrieveUser(username string) (user master.User, err error)
	UpdateUser(id uuid.UUID, user dto.UserUpdate) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn        *gorm.DB
	userService master.UserModelAction
}

func InitializeServices() (err error) {
	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword),
	), &gorm.Config{})
	if err != nil {
		log.Println("[INIT] failed connecting to PostgreSQL")
		return
	}
	log.Println("[INIT] connected to PostgreSQL")

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn:        db,
			userService: master.NewUserAction(db),
		},
	}
	return
}
