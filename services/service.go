package services

import (
	"ip-checker/adapters"
)

// ValidateIP - check if ip address is within the whitelisted countries
func ValidateIP(ipAddress string, countryList []string) bool {
	// get country for ip address
	countryCode, err := adapters.GetCountryCode(ipAddress)
	if err != nil {
		return false
	}

	// check if ip address country is present in whitelisted countries
	for _, eachCountry := range countryList {
		if eachCountry == countryCode {
			return true
		}
	}

	return false
}
