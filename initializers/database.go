package initializers

import (
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	
)

var DB *gorm.DB

func ConnectoDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB,err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=nil {
		log.Fatal("Error connecting to DB")
	}

}