package server

import (
	"fmt"
	"net/http"

	"ypeskov/qr-generator/assets"
	"ypeskov/qr-generator/internal/config"
	"ypeskov/qr-generator/internal/logger"
	mymiddleware "ypeskov/qr-generator/internal/middleware"
	"ypeskov/qr-generator/internal/routes"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Port int
}

func New(cfg *config.Config, log *logger.Logger) *http.Server {
	e := echo.New()
	e.Use(mymiddleware.LoggerMiddleware(log))
	// e.Use(middleware.Logger())

	fileServer := http.FileServer(http.FS(assets.Files))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets", fileServer)))

	routes.RegisterRoutes(e)

	log.Infof("Starting server on port %s", cfg.Port)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: e,
	}
}
