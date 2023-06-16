package handlers

import (
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
	utility.RespondOkayJson(w, response)
}
