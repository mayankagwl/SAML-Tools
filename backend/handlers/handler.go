package handlers

import (
	"encoding/json"
	"github.com/Compile7/SAML-Tools/appErr"
	"github.com/Compile7/SAML-Tools/utility"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetRoutes(router *mux.Router) {
	router.HandleFunc("/", h.RootHandler).Methods(http.MethodGet)
	router.HandleFunc("/post", h.PostHandler).Methods(http.MethodPost)

	//Encode-Decode
	router.HandleFunc("/base64/encode", h.Encode).Methods(http.MethodGet)
	router.HandleFunc("/base64/decode", h.Encode).Methods(http.MethodGet)
	router.HandleFunc("/url/encode", h.Encode).Methods(http.MethodGet)
	router.HandleFunc("/url/decode", h.Encode).Methods(http.MethodGet)
	router.HandleFunc("/base64/inflate", h.Encode).Methods(http.MethodGet)
	router.HandleFunc("/base64/deflate", h.Encode).Methods(http.MethodGet)
}

func (h *Handler) RootHandler(w http.ResponseWriter, _ *http.Request) {
	response := map[string]bool{
		"Ok": true,
	}
	utility.RespondJson(w, http.StatusOK, response)
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Context().Value("body").([]byte)
	if !utility.IsValidJson(body) {
		utility.RespondJson(w, http.StatusBadRequest, appErr.NewError("invalid JSON", "invalid JSON"))
		return
	}
	var bodyMap map[string]string
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		utility.RespondJson(w, http.StatusBadRequest, appErr.NewError("bad format JSON", "bad format JSON"))
		return
	}
	utility.RespondJson(w, http.StatusOK, bodyMap)
}
