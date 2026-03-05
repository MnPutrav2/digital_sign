package pkg

import (
	"encoding/base64"
	"time"
)

func MetaPdf(sign []byte) map[string]string {
	return map[string]string{
		"Title":     "ttd digital",
		"Author":    "saya sendiri",
		"Signature": base64.StdEncoding.EncodeToString(sign),
		"SignerID":  "user-001",
		"Timestamp": time.Now().UTC().String(),
	}
}
