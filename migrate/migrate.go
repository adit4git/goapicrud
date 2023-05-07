package main

import (
	"github.com/adit4git/go-crud/initializers"
	"github.com/adit4git/go-crud/models"
)


func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectoDB()
}


func main() {
 initializers.DB.AutoMigrate(&models.Post{})
}


// Migrate the schema
