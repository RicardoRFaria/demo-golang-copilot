package web

import (
	"demo-golang/internal/services"
	"demo-golang/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Router struct {
	userService services.UserService
}

func NewUserRouter(userService services.UserService) *Router {
	return &Router{userService: userService}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	e.GET("/users/:id", r.getUser)
	e.POST("/users", r.saveUser)
	e.GET("/users", r.listUsers)
}

func (r *Router) getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	context := c.Request().Context()
	user, err := r.userService.GetUser(context, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (r *Router) saveUser(c echo.Context) error {
	context := c.Request().Context()
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := r.userService.SaveUser(context, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User saved successfully")
}

func (r *Router) listUsers(c echo.Context) error {
	context := c.Request().Context()
	users, err := r.userService.ListUsers(context)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
