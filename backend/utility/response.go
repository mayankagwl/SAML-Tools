package utility

import (
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, status int, data interface{}) {
	b, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(b)
}
