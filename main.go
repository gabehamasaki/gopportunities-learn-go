package main

import (
	"github.com/gabehamasaki/gopportunities-learn-go/config"
	"github.com/gabehamasaki/gopportunities-learn-go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v",err.Error())
		return
	}

	// Initialize Router
	router.Initialize()
}
 