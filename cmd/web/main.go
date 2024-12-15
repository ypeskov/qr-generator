package main

import (
	"fmt"

	"ypeskov/qr-generator/internal/config"
	"ypeskov/qr-generator/internal/logger"
	"ypeskov/qr-generator/internal/server"
)

func main() {
	fmt.Println("Hello, world!")
	cfg := config.New()

	logger := logger.New(cfg)

	server := server.New(cfg, logger)

	err := server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
