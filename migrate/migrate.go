package main

import (
	"example.com/authentication/initializers"
	"example.com/authentication/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
	)
}
