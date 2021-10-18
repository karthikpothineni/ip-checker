package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"ip-checker/controllers"
)

// NewRouter - creates echo instance for routing requests
func NewRouter() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.OPTIONS, echo.POST},
	}))

	// endpoint for health check
	router.GET("/status", controllers.GetStatus)

	// endpoint for ip validation
	ipService := router.Group("/")
	var version1 *echo.Group
	version1 = ipService.Group("v1")
	version1.POST("/validate-ip/", controllers.ValidateIP)

	return router
}
