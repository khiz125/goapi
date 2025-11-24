package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (resw *resLoggingWriter) WriteHeader(code int) {
	resw.code = code
	resw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.RequestURI, req.Method)

		reslogwrite := NewResLoggingWriter(w)

		next.ServeHTTP(reslogwrite, req)

		log.Println("response: ", reslogwrite.code)
	})
}
