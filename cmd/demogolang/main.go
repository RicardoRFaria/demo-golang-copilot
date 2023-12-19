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
	companyRepository := repository.NewCompanyRepository()

	userService := services.NewUserService(userRepository)
	companyService := services.NewCompanyService(companyRepository)

	userRoutes := web.NewUserRouter(userService)
	userRoutes.RegisterRoutes(e)
	companyRoutes := web.NewCompanyRouter(companyService)
	companyRoutes.RegisterRoutes(e)

	err := e.Start(":8080")
	if err != nil {
		panic(fmt.Sprintln("error starting server: %w", err))
	}
}
