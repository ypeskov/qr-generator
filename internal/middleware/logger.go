package middleware

import (
	"fmt"
	"time"
	"ypeskov/qr-generator/internal/logger"

	"github.com/labstack/echo/v4"
)

const LoggerKey = "logger"

func LoggerMiddleware(log *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(LoggerKey, log)

			start := time.Now()

			err := next(c)

			status := c.Response().Status
			if err != nil {
				if httpErr, ok := err.(*echo.HTTPError); ok {
					status = httpErr.Code
				} else {
					status = 500
				}
			}

			method := c.Request().Method
			uri := c.Request().RequestURI
			latency := time.Since(start)

			statusColor := getStatusColor(status)
			resetColor := "\033[0m"

			// Формируем сообщение
			message := fmt.Sprintf("%stime=%s, method=%s, uri=%s, status=%s%d%s, latency=%s",
				statusColor,
				time.Now().Format(time.RFC3339),
				method,
				uri,
				statusColor,
				status,
				resetColor,
				latency,
			)

			if status >= 400 {
				log.Errorf("%s", message)
			} else {
				log.Infof("%s", message)
			}

			return err
		}
	}
}

func getStatusColor(status int) string {
	switch {
	case status >= 400 && status <= 599:
		return "\033[31m" // Red
	case status >= 300 && status <= 399:
		return "\033[33m" // Yellow
	default:
		return "" // Regular
	}
}
