package bootstrap

import (
	"fmt"

	"github.com/AlfredoPrograma/diployable/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunApp() {
	env := LoadEnv()
	db := ConnectDB(env.DbConnectionString)

	// Initialize services
	usersService := users.NewService(db)
	// Initialize handlers
	usersHandler := users.NewHandler(usersService)

	e := echo.New()
	e.Use(middleware.AddTrailingSlash())

	apiGroup := e.Group("/api/v1")
	usersGroup := apiGroup.Group("/users")

	usersGroup.GET("/:email", usersHandler.GetUserByEmail)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", env.ApiPort)))
}
