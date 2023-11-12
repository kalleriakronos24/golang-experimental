package migrations

import (
	"fmt"

	database "github.com/kalleriakronos24/mygoapp2nd/db"
	"github.com/kalleriakronos24/mygoapp2nd/models"
)

func Migrate() {

	// auto migration all models
	if err := database.GetDatabaseConnection().AutoMigrate(models.User{}, models.UserLog{}); err != nil {
		panic(err)
	}

	fmt.Println("Database Migrated")
}
