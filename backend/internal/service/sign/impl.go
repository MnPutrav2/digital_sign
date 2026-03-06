package signService

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"os"

	signRepository "github.com/MnPutrav2/go_architecture/internal/repository/sign"
	"github.com/MnPutrav2/go_architecture/pkg"
	"github.com/MnPutrav2/go_architecture/pkg/signature"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type signService struct {
	repo signRepository.SignRepository
}

type SignService interface {
	Signer(ctx context.Context, file multipart.File, x, y string, private *rsa.PrivateKey, public *rsa.PublicKey) (bytes.Buffer, error)
}

func InitSignService(repo signRepository.SignRepository) SignService {
	return &signService{repo: repo}
}

// Entry

func (s *signService) Signer(ctx context.Context, file multipart.File, x, y string, private *rsa.PrivateKey, public *rsa.PublicKey) (bytes.Buffer, error) {
	// Generate qr code
	qrBytes, err := pkg.GenerateQR("https://yourdomain.com/verify/123")
	if err != nil {
		return bytes.Buffer{}, err
	}

	if err = os.WriteFile("temp_qr.png", qrBytes, 0644); err != nil {
		return bytes.Buffer{}, err
	}
	defer os.Remove("temp_qr.png")

	// Append to pdf
	var buf bytes.Buffer
	alg := fmt.Sprintf("pos:tl, off:%s %s, scale:0.25 abs, rot:0", x, y)
	wm, err := pdfcpu.ParseImageWatermarkDetails("temp_qr.png", alg, true, types.POINTS)
	if err != nil {
		return bytes.Buffer{}, err
	}

	if err := api.AddProperties(file, &buf, pkg.MetaPdf([]byte("ke")), nil); err != nil {
		return bytes.Buffer{}, err
	}

	var qg bytes.Buffer
	if err := api.AddWatermarks(bytes.NewReader(buf.Bytes()), &qg, nil, wm, nil); err != nil {
		return bytes.Buffer{}, err
	}

	// Create signature
	hash := qg.Bytes()

	sign, err := signature.SignPDF(hash, private)
	if err != nil {
		return bytes.Buffer{}, err
	}

	if err := s.repo.InsertSignature(ctx, base64.StdEncoding.EncodeToString(sign)); err != nil {
		return bytes.Buffer{}, err
	}

	fmt.Println("SIZE SIGN:", len(hash))

	fmt.Println("signature: ", base64.StdEncoding.EncodeToString(sign))
	return qg, nil
}
