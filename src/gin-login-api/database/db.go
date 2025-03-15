package auth_database

import (
	"fmt"
	auth_models "go-backend/src/gin-login-api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// connectDatabase initializes the database connection
func ConnectDatabase() {

	var err error
	DB, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ,", err)
	}

	// Auto-migrate User Model
	err = DB.AutoMigrate((&auth_models.User{}))

	if err!= nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("->> Database connected successfully!")

}
