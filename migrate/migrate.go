package main

import (
	"github/com/gamxu/go-api-tutorial/initializers"
	"github/com/gamxu/go-api-tutorial/models"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}