package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"ip-checker/controllers"
)

// NewRouter for routing requests
func NewRouter() *echo.Echo {

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.OPTIONS, echo.POST},
	}))

	// Endpoints for healthcheck
	router.GET("/status", controllers.GetStatus)

	ipService := router.Group("/")
	// this group is for read APIS
	var version1 *echo.Group
	version1 = ipService.Group("v1")
	version1.POST("/validate-ip/", controllers.ValidateIP)

	return router
}
