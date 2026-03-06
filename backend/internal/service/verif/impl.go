package verifService

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"io"
	"mime/multipart"

	verifRepository "github.com/MnPutrav2/go_architecture/internal/repository/verif"
	"github.com/MnPutrav2/go_architecture/pkg/signature"
)

type verifService struct {
	repo verifRepository.VerifRepository
}

type VerifService interface {
	Verif(ctx context.Context, file multipart.File, public *rsa.PublicKey, sign string) error
}

func InitVerifService(repo verifRepository.VerifRepository) VerifService {
	return &verifService{repo: repo}
}

// Entry
func (s *verifService) Verif(ctx context.Context, file multipart.File, public *rsa.PublicKey, sign string) error {
	if seeker, ok := file.(io.Seeker); ok {
		seeker.Seek(0, 0)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	dec, _ := base64.StdEncoding.DecodeString(sign)
	if err := signature.VerifyPDF(data, dec, public); err != nil {
		return err
	}

	return nil
}
