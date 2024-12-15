package render

import (
	"ypeskov/qr-generator/internal/logger"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component, log logger.Logger) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		log.Errorf("Error rendering component: %e\n", err)
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
