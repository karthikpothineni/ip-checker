package services

import (
	"github.com/stretchr/testify/suite"
	"ip-checker/adapters"
	"ip-checker/config"
	"ip-checker/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ServiceTestSuite - test suite for validating service functions
type ServiceTestSuite struct {
	suite.Suite
}

// SetupTest - This will run before all the tests in the suite
func (s *ServiceTestSuite) SetupSuite() {
	logger.Init()
	config.Init("../config/")
	_ = adapters.InitGeoIPReader("../" + config.GetConfig().GetString("geoip.filepath"))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestIPTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

// TestIPWhitelistWithValidIP - tests if IP is whitelisted when IP Address is valid
func (s *ServiceTestSuite) TestIPWhitelistWithValidIP() {
	check := assert.New(s.T())

	isWhitelisted := ValidateIP("206.71.50.230", []string{"US"})
	check.Equal(true, isWhitelisted)
}

// TestIPWhitelistWithInvalidIP - tests if IP is whitelisted when IP Address is invalid
func (s *ServiceTestSuite) TestIPWhitelistWithInvalidIP() {
	check := assert.New(s.T())

	isWhitelisted := ValidateIP("206.71.50.", []string{"US"})
	check.Equal(false, isWhitelisted)
}

// TestIPWhitelistWithEmptyIP - tests if IP is whitelisted when IP Address is empty
func (s *ServiceTestSuite) TestIPWhitelistWithEmptyIP() {
	check := assert.New(s.T())

	isWhitelisted := ValidateIP("", []string{"IN"})
	check.Equal(false, isWhitelisted)
}

// TestIPWhitelistWithEmptyCountryList - tests if IP is whitelisted when country list is empty
func (s *ServiceTestSuite) TestIPWhitelistWithEmptyCountryList() {
	check := assert.New(s.T())

	isWhitelisted := ValidateIP("206.71.50.230", []string{})
	check.Equal(false, isWhitelisted)
}

// TestIPWhitelistWithEmptyIPAndCountryList - tests if IP is whitelisted when IP Address and country list is empty
func (s *ServiceTestSuite) TestIPWhitelistWithEmptyIPAndCountryList() {
	check := assert.New(s.T())

	isWhitelisted := ValidateIP("", []string{})
	check.Equal(false, isWhitelisted)
}
