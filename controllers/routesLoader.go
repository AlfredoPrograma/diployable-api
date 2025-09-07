package controllers

import "github.com/labstack/echo/v4"

// RoutesLoader defines an interface for loading routes into an Echo group.
// Implementations of this interface should provide the logic to register
// HTTP routes on the provided *echo.Group.
type routesLoader interface {
	LoadRoutes(group *echo.Group)
}
