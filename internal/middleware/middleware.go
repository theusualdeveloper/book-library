package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
				log.Printf("recovered from panic: %v", p)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		start := time.Now()
		next.ServeHTTP(rw, r)
		reqTimer := time.Since(start)
		reqTime := start.Format("2006/01/02 15:04:05")
		fmt.Printf("%s %s %s %d %v\n",
			reqTime,
			r.Method,
			r.URL.Path,
			rw.statusCode,
			reqTimer,
		)
	})
}
