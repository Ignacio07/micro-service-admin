package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/Ignacio07/micro-service-admin/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server starting...")
	godotenv.Load()
	fmt.Println("Loaded env variables...")

	// Connect to database
	db.Connect()

	//create fiber app
	app := fiber.New()
	app.Use(cors.New())

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not set in .env file")
	}

}
