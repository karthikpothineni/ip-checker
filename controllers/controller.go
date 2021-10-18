package controllers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"ip-checker/models"
	"ip-checker/services"
	"ip-checker/utils"
)

// GetStatus - Controller function to check the status of the rest server
func GetStatus(c echo.Context) error {
	response := utils.BuildResponse("success", "", "Server is UP and Running")
	return c.JSON(http.StatusOK, response)
}

// ValidateIP - Controller function to validate ip address with the help of GeoIP database
func ValidateIP(c echo.Context) error {

	// validate request
	validationRequest := models.IPValidationRequest{}
	if err := c.Bind(&validationRequest); err != nil {
		response := utils.BuildResponse("fail", "Invalid request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if strings.TrimSpace(validationRequest.IPAddress) == "" {
		response := utils.BuildResponse("fail", "IP address cannot be empty", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if len(validationRequest.CountyWhiteList) == 0 {
		response := utils.BuildResponse("fail", "country list cannot be empty", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	// check ip address
	isValidIP := services.ValidateIP(validationRequest.IPAddress, validationRequest.CountyWhiteList)

	response := utils.BuildResponse("success", "", isValidIP)
	return c.JSON(http.StatusOK, response)
}
