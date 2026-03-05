package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"hash"
)

func Generate(h hash.Hash) ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA256,
		h.Sum(nil)[:],
	)

	return signature, nil
}
