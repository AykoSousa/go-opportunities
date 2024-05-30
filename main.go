package main

import (
	"github.com/AykoSousa/go-opportunities/config"
	"github.com/AykoSousa/go-opportunities/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialiation error: %v", err)
		return
	}
	// Initialize routes
	router.Initialize()
}