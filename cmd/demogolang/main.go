package main

import (
	"demo-golang/internal/repository"
	"demo-golang/internal/services"
	"demo-golang/web"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)

	router := web.NewUserRouter(userService)
	router.RegisterRoutes(e)

	err := e.Start(":8080")
	if err != nil {
		panic(fmt.Sprintln("error starting server: %w", err))
	}
}
