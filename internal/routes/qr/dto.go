package qr

import "fmt"

type GenerateQRCodeRequest struct {
	Type    string `form:"type" validate:"required"`
	Content string `form:"content" validate:"required"`
}

func (r GenerateQRCodeRequest) String() string {
	return fmt.Sprintf("Type: %s, Content: %s", r.Type, r.Content)
}
