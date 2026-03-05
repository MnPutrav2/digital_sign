package response

import (
	"bytes"
	"net/http"

	logging "github.com/MnPutrav2/go_architecture/pkg/log"
)

func File(file bytes.Buffer, log string, ty string, w http.ResponseWriter, r *http.Request) {
	logging.Log(log, ty, r)
	w.Header().Set("Content-Disposition", "attachment; filename=signed.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(file.Bytes())
}
