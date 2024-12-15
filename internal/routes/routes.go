package routes

import (
	"ypeskov/qr-generator/internal/routes/health"
	"ypeskov/qr-generator/internal/routes/home"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", home.Home)

	healthGroup := e.Group("/health")
	health.RegisterRoutes(healthGroup)
}
