package qr

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"

	"ypeskov/qr-generator/internal/logger"
	"ypeskov/qr-generator/internal/middleware"
	"ypeskov/qr-generator/internal/render"
	"ypeskov/qr-generator/templates/components"
)

func GenerateQRCode(c echo.Context) error {
	log := c.Get(middleware.LoggerKey).(*logger.Logger)
	log.Info("Generating QR code")

	var dataToCode GenerateQRCodeRequest
	if err := c.Bind(&dataToCode); err != nil {
		log.Error("Failed to bind request", err)
		return render.Render(c, http.StatusBadRequest, nil, *log)
	}
	log.Info(fmt.Sprintf("Data to encode: {%+v} ", dataToCode))

	qrCode, err := qrcode.Encode(getQRCodeData(dataToCode), qrcode.Medium, 256)
	if err != nil {
		log.Error("Failed to generate QR code", err)
		return render.Render(c, http.StatusInternalServerError, nil, *log)
	}

	qrCodeBase64 := base64.StdEncoding.EncodeToString(qrCode)
	img := fmt.Sprintf("data:image/png;base64,%s", qrCodeBase64)
	component := components.QrCode(img)

	return render.Render(c, http.StatusOK, component, *log)
}

func getQRCodeData(dataToCode GenerateQRCodeRequest) string {
	return fmt.Sprintf("%s: %s", dataToCode.Type, dataToCode.Content)
}
