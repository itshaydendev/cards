package main

import (
	"github.com/itshaydendev/cards/internal/api"
	"github.com/itshaydendev/cards/pkg/logger"
)

func main() {
	logger.Info("Starting Cards...")
	api.StartServer()
}
