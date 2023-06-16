package handlers

import (
	"github.com/Compile7/SAML-Tools/utility"
	"net/http"
)

func (h *Handler) Encode(w http.ResponseWriter, _ *http.Request) {
	response := map[string]string{
		"Status": "ok",
	}
	utility.RespondJson(w, http.StatusOK, response)
}
