package migrations

import (
	"fmt"

	database "github.com/kalleriakronos24/mygoapp2nd/db"
	"github.com/kalleriakronos24/mygoapp2nd/models"
	masterModel "github.com/kalleriakronos24/mygoapp2nd/models/master"
)

func Migrate() {
	// auto migration all models
	if err := database.GetDatabaseConnection().AutoMigrate(
		// master
		masterModel.User{},
		masterModel.Module{},
		masterModel.SysOptions{},
		// common
		models.UserLog{},
		models.UserModule{},
		models.UserModulePermission{},
	); err != nil {
		panic(err)
	}

	fmt.Println("Database Migrated")
}
