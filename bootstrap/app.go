package bootstrap

import (
	"fmt"

	"github.com/AlfredoPrograma/diployable/controllers"
	"github.com/AlfredoPrograma/diployable/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunApp() {
	env := LoadEnv()
	db := ConnectDB(env.DbConnectionString)

	// Initialize services
	usersService := services.NewUsersService(db)
	// Initialize handlers
	usersController := controllers.NewUsersController(usersService)

	e := echo.New()
	e.Use(middleware.AddTrailingSlash())

	apiGroup := e.Group("/api/v1")

	usersController.LoadRoutes(apiGroup.Group("/users"))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", env.ApiPort)))
}
