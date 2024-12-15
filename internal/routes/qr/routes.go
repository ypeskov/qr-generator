package qr

import "github.com/labstack/echo/v4"

func RegisterRoutes(group *echo.Group) {
	group.POST("/generate", GenerateQRCode)
}
