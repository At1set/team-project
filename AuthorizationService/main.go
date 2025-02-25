package main

import (
	"authorization-service/handlers"
	"authorization-service/middleware"
	"authorization-service/storage"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "1324"
	dbname   = "AuthUsersDb"
)

func main() {
	app := fiber.New()

	// Инициализация хранилища

	authStorage, err := storage.NewAuthStorage(host, port, user, password, dbname)

	if err != nil {
		fmt.Println("Не удалось подключится к базе данных")
		fmt.Println(err)
	}
	authHandler := handlers.NewAuthHandler(authStorage)

	app.Static("/", "static")
	app.Post("/login", authHandler.Login)

	authorizedGroup := app.Group("/api/admin")
	authorizedGroup.Use(middleware.JWTAdminMiddleware())
	authorizedGroup.Post("/users", authHandler.Register)
	authorizedGroup.Get("/users", authHandler.GetUsers)
	authorizedGroup.Get("/users/:id", authHandler.GetUser)
	authorizedGroup.Delete("/users/:id", authHandler.DeleteUser)

	app.Listen(":8080")
}
