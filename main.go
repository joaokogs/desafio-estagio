package main

import (
	"github.com/joaokogs/desafio-estagio/config"
	"github.com/joaokogs/desafio-estagio/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Init config

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Init router
	router.Init()
}
