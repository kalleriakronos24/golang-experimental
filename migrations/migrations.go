package migrations

import (
	"fmt"

	database "github.com/kalleriakronos24/mygoapp2nd/db"
	"github.com/kalleriakronos24/mygoapp2nd/models"
	masterModels "github.com/kalleriakronos24/mygoapp2nd/models/master"
)

func Migrate() {
	// auto migration all models
	if err := database.GetDatabaseConnection().AutoMigrate(
		// master
		masterModels.User{},
		masterModels.Module{},
		masterModels.SysOptions{},
		// common
		models.UserLog{},
		models.UserModule{},
		models.UserModulePermission{},
	); err != nil {
		panic(err)
	}

	fmt.Println("Database Migrated")
}
