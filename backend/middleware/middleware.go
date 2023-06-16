package middleware

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

const JSONContentType = "application/json"

type GoMiddleware struct {
}

func NewMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

func (g *GoMiddleware) PostBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyPayload := make([]byte, 0)
		if r.Method != http.MethodGet {
			contentType := strings.ToLower(r.Header.Get("Content-Type"))
			if strings.HasPrefix(contentType, JSONContentType) {
				b, err := ioutil.ReadAll(r.Body)
				if err == nil && len(b) > 0 {
					bodyPayload = b
				}
			}
		}
		ctx := context.WithValue(r.Context(), "body", bodyPayload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
