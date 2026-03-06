package response

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/go_architecture/internal/model"
	logging "github.com/MnPutrav2/go_architecture/pkg/log"
)

func Message(message string, log string, ty string, code int, w http.ResponseWriter, r *http.Request) {
	var status string

	switch ty {
	case "INFO":
		status = "Success"
	case "WARN":
		status = "Failed"
	case "ERROR":
		status = "Error"
	}

	res, _ := json.Marshal(model.ResponseMessage{Message: message, Meta: model.Meta{
		Code:   code,
		Status: status,
	}})
	logging.Log(log, ty, r)
	w.WriteHeader(code)
	w.Write(res)
}
