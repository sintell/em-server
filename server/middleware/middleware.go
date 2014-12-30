package middleware

import (
	"log"
	"net/http"
)

func NewLogging(l *log.Logger) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Printf("%s\t%s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}

}
