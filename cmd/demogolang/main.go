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
	companyRoutes := web.NewCompanyRouter(companyService)
	employeeStateRoutes := web.NewEmployeeStateRouter()

	userRoutes.RegisterRoutes(e)
	companyRoutes.RegisterRoutes(e)
	employeeStateRoutes.RegisterRoutes(e)

	err := e.Start(":8080")
	if err != nil {
		panic(fmt.Sprintln("error starting server: %w", err))
	}
}
