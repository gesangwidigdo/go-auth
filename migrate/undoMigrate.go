package main

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.Migrator().DropTable(&models.User{})
}
