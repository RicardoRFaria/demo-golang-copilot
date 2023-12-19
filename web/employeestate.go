package web

import (
	"demo-golang/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EmployeeStateRouter struct {
}

func NewEmployeeStateRouter() *EmployeeStateRouter {
	return &EmployeeStateRouter{}
}

func (r *EmployeeStateRouter) RegisterRoutes(e *echo.Echo) {
	e.GET("/employee/", r.GetEmployeeState)
}

func (r *EmployeeStateRouter) GetEmployeeState(c echo.Context) error {
	// For simplicity, let's use a static EmployeeState
	employee := model.EmployeeStateUnaligned{
		IsActive:  true,
		Age:       30,
		Name:      "John Doe",
		IsMarried: false,
	}

	// Respond with the EmployeeState as JSON
	return c.JSON(http.StatusOK, employee)
}
