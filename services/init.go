package services

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/mygoapp2nd/config"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
	masterModels "github.com/kalleriakronos24/mygoapp2nd/models/master"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Handler HandlerFunc

type HandlerFunc interface {
	// AuthenticateUser ====== Authentication Handlers ====== //
	AuthenticateUser(p dto.UserLogin) (token string, err error)

	// RegisterUser ====== User Handlers ====== //
	RegisterUser(p dto.UserSignup) (err error)
	RetrieveUser(username string) (m masterModels.User, err error)
	UpdateUser(id uuid.UUID, p dto.UserUpdate) (err error)

	// CreateMasterModule ====== Master Module Handlers ====== //
	CreateMasterModule(p dto.CreateMasterModule) (m masterModels.Module, err error)
	UpdateMasterModule(id uuid.UUID, p dto.UpdateMasterModule) (err error)
	RetrieveMasterModule(p dto.RetrieveOneMasterModule) (m masterModels.Module, err error)
	DeleteMasterModule(id uuid.UUID) (err error)
	CheckExistingMasterModule(moduleName string, id string) (err error)

	// ====== Master System Option Handlers ====== //
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn                    *gorm.DB
	userModel               masterModels.UserModelAction
	masterModuleModel       masterModels.ModuleModelAction
	masterSystemOptionModel masterModels.SysOptionsModelAction
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
		log.Println("[INIT] failed connecting to PostgresSQL")
		return
	}
	log.Println("[INIT] connected to PostgresSQL")

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn:                    db,
			userModel:               masterModels.NewUserAction(db),
			masterModuleModel:       masterModels.NewModuleAction(db),
			masterSystemOptionModel: masterModels.NewSysOptionsAction(db),
		},
	}
	return
}
