package models

// IPValidationRequest - represents ip validation request details
type IPValidationRequest struct {
	IPAddress       string   `json:"ip_address"`
	CountyWhiteList []string `json:"country_whitelist"`
}
