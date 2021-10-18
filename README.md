ip-checker
==========

This application is responsible for checking if country of an IP Address is part of the whitelisted countries or not.

## Setup
### Local
1. Clone the repository under GOPATH
2. Install dependencies using ```go mod download```
3. Run the application using ```go run main.go```
4. By default, server starts on the port 5000. If required, one can change the port in config/config.toml and docker-compose.yml
5. Stop the application by pressing CTRL+C
### Docker
Docker compose internally runs the linter, tests before building the application. If there is any error with linter or tests, build will be failed. Run docker compose using 

```docker-compose up```
### Run Linter
```golangci-lint run -v -c golangci.yml```
### Run Tests
```go test -v -cover ./...```

I have added the tests for **adapters**,**services**,**utils**. Code coverage is greater than **85%**

## API Usage
### 1) 
### Sample Request(Valid):
**HTTP Method:** POST

**URL:** http://localhost:5000/v1/validate-ip/

**Request Body:**
```json
{
    "ip_address": "206.71.50.230",
    "country_whitelist": ["IN","US"]
}
```
### Sample Response(Valid):
**Status:** 200 OK

**Response Body:**
```json
{
    "data": true,
    "status": "success"
}
```

### 2)
### Sample Request(Invalid):
**HTTP Method:** POST

**URL:** http://localhost:5000/v1/validate-ip/

**Request Body:**
```json
{
    "ip_address": "",
    "country_whitelist": ["IN","US"]
}
```
### Sample Response(Invalid):
**Status:** 400 Bad Request

**Response Body:**
```json
{
    "error": "IP address cannot be empty",
    "status": "fail"
}
```

## Plan for keeping the mapping data up to date
1) Create a Rule using Amazon EventBridge which internally triggers a Lambda function periodically.
2) Lamba function internally calls the Update Mapping Data API in our service. 
3) This API internally downloads the latest mapping data and re-initializes the GeoIP Reader Object.
