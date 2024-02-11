package middleware

import (
	"net/http"
)

const (
	MB = 1 << 20
)

func ImageSizeValidator(errorHandlerFunc func(w http.ResponseWriter)) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, multipertFileHeader, err := r.FormFile("file")
			if err != nil {
				errorHandlerFunc(w)
				return
			}

			if multipertFileHeader.Size > (5 * MB) {
				errorHandlerFunc(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
