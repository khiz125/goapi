package middlewares

import (
	"log"
	"net/http"

	"github.com/khiz125/goapi/common"
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
		traceID := newTraceID()

		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := common.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)
		reslogwrite := NewResLoggingWriter(w)

		next.ServeHTTP(reslogwrite, req)

		log.Printf("[%d]response: %d", traceID, reslogwrite.code)
	})
}
