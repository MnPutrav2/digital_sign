package response

import (
	"encoding/json"
	"net/http"

	"github.com/MnPutrav2/go_architecture/internal/model"
	logging "github.com/MnPutrav2/go_architecture/pkg/log"
)

func Body(response any, param []string, log string, ty string, w http.ResponseWriter, r *http.Request) {
	var status string

	switch ty {
	case "INFO":
		status = "Success"
	case "WARN":
		status = "Failed"
	case "ERROR":
		status = "Error"
	}

	res, _ := json.Marshal(model.ResponseBody{Response: response, Meta: model.Meta{
		Code:   200,
		Status: status,
	}})
	logging.Log(log, ty, r)
	w.WriteHeader(200)
	w.Write(res)
}
