package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jared0o/auth-microservice/controllers"
	"github.com/jared0o/auth-microservice/initializers"
	"github.com/jared0o/auth-microservice/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	app := fiber.New()

	app.Post("/singup", controllers.Signup)
	app.Post("/login", controllers.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))
	app.Get("/validate", middleware.RequireAuth, controllers.Validate)

	app.Listen(os.Getenv("PORT"))
}
