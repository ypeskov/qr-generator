package main

import (
	"ypeskov/qr-generator/internal/config"
	"ypeskov/qr-generator/internal/logger"
	"ypeskov/qr-generator/internal/server"
)

func main() {
	cfg := config.New()

	logger := logger.New(cfg)

	server := server.New(cfg, logger)

	err := server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
