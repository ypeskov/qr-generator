package server

import (
	"fmt"
	"net/http"
	"ypeskov/qr-generator/internal/config"
	"ypeskov/qr-generator/internal/logger"
	"ypeskov/qr-generator/internal/middleware"
	"ypeskov/qr-generator/internal/routes"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Port int
}

func New(cfg *config.Config, logger *logger.Logger) *http.Server {
	echo := echo.New()
	echo.Use(middleware.LoggerMiddleware(logger))

	routes.RegisterRoutes(echo)

	logger.Infof("Starting server on port %s", cfg.Port)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: echo,
	}
}
