package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"ip-checker/config"
	"ip-checker/logger"
)

// GeoIPTestSuite - test suite for ip validation
type GeoIPTestSuite struct {
	suite.Suite
	geoIPPath string
}

// SetupTest - This will run before all the tests in the suite
func (s *GeoIPTestSuite) SetupSuite() {
	logger.Init()
	config.Init("../config/")
	s.geoIPPath = "../" + config.GetConfig().GetString("geoip.filepath")
	_ = InitGeoIPReader(s.geoIPPath)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestIPTestSuite(t *testing.T) {
	suite.Run(t, new(GeoIPTestSuite))
}

// TestInitGeoIPReader - tests when geoip reader is initialized
func (s *GeoIPTestSuite) TestInitGeoIPReader() {
	check := assert.New(s.T())

	err := InitGeoIPReader(s.geoIPPath)
	check.Nil(err)
	check.NotNil(GeoIP)
}

// TestGetCountryCodeValidIP - tests when country code is retrieved for a valid IP
func (s *GeoIPTestSuite) TestGetCountryCodeValidIP() {
	check := assert.New(s.T())

	countryCode, err := GetCountryCode("206.71.50.230")
	check.Equal("US", countryCode)
	check.Nil(err)
}

// TestGetCountryCodeInValidIP - tests when country code is retrieved for a invalid IP
func (s *GeoIPTestSuite) TestGetCountryCodeInValidIP() {
	check := assert.New(s.T())

	countryCode, err := GetCountryCode("206.71.50.")
	check.Equal("", countryCode)
	check.NotNil(err)
}

// TestGetCountryCodeEmptyIP - tests when country code is retrieved for a empty IP
func (s *GeoIPTestSuite) TestGetCountryCodeEmptyIP() {
	check := assert.New(s.T())

	countryCode, err := GetCountryCode("0.0.0.0")
	check.Equal("", countryCode)
	check.Equal("empty country code", err.Error())
}
