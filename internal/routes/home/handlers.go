package home

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"ypeskov/qr-generator/internal/logger"
	"ypeskov/qr-generator/internal/middleware"
	"ypeskov/qr-generator/internal/render"
	"ypeskov/qr-generator/templates/components"
)

func Home(c echo.Context) error {
	log := c.Get(middleware.LoggerKey).(*logger.Logger)
	component := components.HomePage()

	return render.Render(c, http.StatusOK, component, *log)
}
