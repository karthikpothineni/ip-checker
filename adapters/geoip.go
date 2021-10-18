package adapters

import (
	"errors"
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"

	"ip-checker/logger"
)

var GeoIP *geoip2.Reader

// InitGeoIPReader - Inits GeoIP
func InitGeoIPReader(path string) (err error) {
	GeoIP, err = geoip2.Open(path)
	if err != nil {
		logger.Log.Errorf("GeoIP Init Error: %s", err.Error())
		return err
	}

	return nil
}

// GetCountryCode - gets country for a given IP
func GetCountryCode(ipAddress string) (string, error) {
	var countryCode string

	// get country details
	ip := net.ParseIP(ipAddress)
	countryObj, err := GeoIP.Country(ip)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"ip":  ipAddress,
			"err": err,
		}).Error("error while retrieving country")
		return countryCode, err
	}

	// check if country is empty
	countryCode = countryObj.Country.IsoCode
	if len(countryCode) == 0 {
		logger.Log.WithFields(logrus.Fields{
			"ip": ipAddress,
		}).Error("empty country code")
		return countryCode, errors.New("empty country code")
	}

	return countryCode, nil
}
