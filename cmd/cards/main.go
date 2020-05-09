package main

import (
	"github.com/itshaydendev/cards/pkg/api"
	"github.com/itshaydendev/cards/pkg/logger"
)

func main() {
	logger.Info("Starting Cards...")
	api.StartServer()
}
