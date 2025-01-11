package main

import (
	"his_apis_go/initializers"
	// "his_apis_go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}
func main() {
	// initializers.DB.AutoMigrate(&models.Task{})
	// initializers.DB.AutoMigrate(&models.Tasks_categorie{})
}
