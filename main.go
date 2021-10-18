package main

import (
	"ip-checker/adapters"
	"ip-checker/config"
	"ip-checker/logger"
	"ip-checker/server"
)

func main() {
	logger.Init()
	config.Init("config/")
	if err := adapters.InitGeoIPReader(config.GetConfig().GetString("geoip.filepath")); err != nil {
		return
	}
	server.Init()
}
