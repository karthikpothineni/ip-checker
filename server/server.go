package server

import (
	"ip-checker/config"
	"ip-checker/logger"
)

// Init - function to initialize the server
func Init() {
	conf := config.GetConfig()
	logger.Log.Info("Initializing Rest server")
	r := NewRouter()
	if err := r.Start(conf.GetString("general.server_port")); err != nil {
		logger.Log.Fatal("Unable to bring service up", err.Error())
	}
}
