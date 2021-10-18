package utils

import "github.com/labstack/echo"

// BuildResponse - creates a response object out of data given
func BuildResponse(status string, error string, data interface{}) echo.Map {
	response := echo.Map{
		"status": status,
	}

	if error != "" {
		response["error"] = error
		return response
	}
	response["data"] = data

	return response
}
