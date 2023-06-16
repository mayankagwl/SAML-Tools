package utility

import (
	"encoding/json"
	"net/http"
)

func RespondOkayJson(w http.ResponseWriter, data interface{}) {
	b, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}
