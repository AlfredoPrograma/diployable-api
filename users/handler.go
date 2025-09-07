package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	GetUserByEmail(c echo.Context) error
}

type handler struct {
	usersService Service
}

// GetUserByEmail implements Handler.
func (h handler) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")
	user, err := h.usersService.GetUserByEmail(c.Request().Context(), email)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func NewHandler(usersService Service) Handler {
	return handler{
		usersService: usersService,
	}
}
