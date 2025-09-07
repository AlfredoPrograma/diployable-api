package controllers

import (
	"net/http"

	"github.com/AlfredoPrograma/diployable/services"
	"github.com/labstack/echo/v4"
)

type UsersController interface {
	GetUserByEmail(c echo.Context) error
	routesLoader
}

type usersController struct {
	usersService services.UsersService
}

func (ctrl usersController) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")
	user, err := ctrl.usersService.GetUserByEmail(c.Request().Context(), email)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (ctrl usersController) LoadRoutes(group *echo.Group) {
	group.GET("/:email", ctrl.GetUserByEmail)
}

func NewUsersController(usersService services.UsersService) UsersController {
	return usersController{
		usersService: usersService,
	}
}
