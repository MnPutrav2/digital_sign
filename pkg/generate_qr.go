package pkg

import (
	"image/color"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(content string) ([]byte, error) {
	qr, err := qrcode.New(content, qrcode.High)
	if err != nil {
		return nil, err
	}

	qr.ForegroundColor = color.Black
	qr.BackgroundColor = color.White

	return qr.PNG(256)
}
