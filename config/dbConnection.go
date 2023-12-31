package config

import (
	"fmt"
	"os"

	"github.com/Ignacio07/micro-service-admin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var dbURL = os.Getenv("DB_URL")
	if dbURL == "" {
		panic("DB_URL environment variable missing")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	autoMigrate(DB)

}

func autoMigrate(connection *gorm.DB) {
	//connection.Debug().AutoMigrate(&models.User{}, &models.Task{}) //Informacion por consola
	connection.AutoMigrate(&models.User{})
}
