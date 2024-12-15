package middleware

import (
	"ypeskov/qr-generator/internal/logger"

	"github.com/labstack/echo/v4"
)

const LoggerKey = "logger"

func LoggerMiddleware(log *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(LoggerKey, log)
			return next(c)
		}
	}
}
