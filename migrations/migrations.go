package migrations

import (
	"fmt"

	database "github.com/kalleriakronos24/mygoapp2nd/db"
	"github.com/kalleriakronos24/mygoapp2nd/models"
)

func Migrate() {

	// auto migration all models
	database.GetDatabaseConnection().AutoMigrate(models.User{}, models.UserLog{})

	fmt.Println("Database Migrated")
}
