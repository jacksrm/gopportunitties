package main

import (
	"fmt"

	"github.com/jacksrm/gopportunitties/config"
	"github.com/jacksrm/gopportunitties/internal/router"
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		fmt.Println()
		return
	}

	router.Initialize()
}
