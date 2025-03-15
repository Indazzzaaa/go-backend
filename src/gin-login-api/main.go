package main

import (
	"fmt"
	auth_database "go-backend/src/gin-login-api/database"
	"go-backend/src/gin-login-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("->> Running services in GIN.")

	// connect to the database
	auth_database.ConnectDatabase()

	router := gin.Default()

	// Register authentication routes
	routes.AuthRoutes(router)

	// start server
	router.Run(":8080")

}
