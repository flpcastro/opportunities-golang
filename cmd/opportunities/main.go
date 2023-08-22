package main

import (
	"github.com/flpcastro/opportunities-api-go/internal/configs"
	"github.com/flpcastro/opportunities-api-go/internal/router"
	"github.com/flpcastro/opportunities-api-go/pkg/database"
)

var (
	logger *configs.Logger
)

func main() {
	logger = configs.GetLogger("main")

	// Init DB
	err := database.Init()
	if err != nil {
		logger.Errorf("init database error: %v", err)
		return
	}

	// Init router
	router.Initialize()
}
