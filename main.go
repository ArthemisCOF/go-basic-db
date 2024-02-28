package main

import (
	"go-fiber/config"
	usercontroller "go-fiber/controller/user"
	"go-fiber/repository"
	"log"

	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := new(config.DBConfig)

	err = env.Parse(dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	repository.ConnectDatabase(dbConfig)
	api := app.Group("/api")
	v1 := api.Group("/v1")
	userGroup := v1.Group("/user")
	userGroup.Post("/register", usercontroller.RegiterUser)

	app.Listen(":8000")
}
