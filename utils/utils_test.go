package utils

import (
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// TestBuildResponseNoError - tests response creation when there is no error
func TestBuildResponseNoError(t *testing.T) {
	check := assert.New(t)
	expectedResponse := echo.Map{
		"status": "success",
		"data":   true,
	}

	actualResponse := BuildResponse("success", "", true)
	check.Equal(expectedResponse, actualResponse)
}

// TestBuildResponseWithError - tests response creation when there is error
func TestBuildResponseWithError(t *testing.T) {
	check := assert.New(t)
	expectedResponse := echo.Map{
		"status": "fail",
		"error":  "invalid request",
	}

	actualResponse := BuildResponse("fail", "invalid request", nil)
	check.Equal(expectedResponse, actualResponse)
}
