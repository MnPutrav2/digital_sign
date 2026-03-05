package route

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MnPutrav2/digital_sign/pkg"
	"github.com/MnPutrav2/digital_sign/pkg/signature"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func Sign() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		defer file.Close()

		x := r.FormValue("x")
		y := r.FormValue("y")

		var buf bytes.Buffer
		h := sha256.New()

		if _, err := io.Copy(h, file); err != nil {
			http.Error(w, "Process failed", http.StatusInternalServerError)
			return
		}

		signature, err := signature.Generate(h)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		qrBytes, err := pkg.GenerateQR("https://yourdomain.com/verify/123")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = os.WriteFile("temp_qr.png", qrBytes, 0644); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer os.Remove("temp_qr.png")

		alg := fmt.Sprintf("pos:bc, off:%s %s, scale:0.2 abs, rot:0", x, y)
		wm, err := pdfcpu.ParseImageWatermarkDetails("temp_qr.png", alg, true, types.POINTS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := api.AddProperties(file, &buf, pkg.MetaPdf(signature), nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var qg bytes.Buffer
		if err := api.AddWatermarks(bytes.NewReader(buf.Bytes()), &qg, nil, wm, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=signed.pdf")
		w.Header().Set("Content-Type", "application/pdf")
		w.Write(qg.Bytes())
	}
}
