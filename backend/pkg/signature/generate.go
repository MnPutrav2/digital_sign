package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func SignPDF(h []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(h)

	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA256,
		hash[:],
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("SIGN HASH: ", hex.EncodeToString(hash[:]))
	return signature, nil
}

func VerifyPDF(h []byte, signature []byte, publicKey *rsa.PublicKey) error {

	hashed := sha256.Sum256(h)

	err := rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA256,
		hashed[:],
		signature,
	)

	fmt.Println("VERIF HASH: ", hex.EncodeToString(hashed[:]))
	return err
}
