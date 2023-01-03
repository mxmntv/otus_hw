package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mxmntv/otus_hw/hw12_calendar/pkg/logger"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func loggingMiddleware(logger logger.Interface, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(w, r)
		lat := time.Since(start)
		logline := fmt.Sprintf("%s [%s] %s %s %s %d %d %s", r.RemoteAddr, start.Format("02/Jan/2006:15:04:05 -0700"),
			r.Method, r.RequestURI, r.Proto, lrw.statusCode, lat, r.UserAgent())
		logger.Info(logline)
	})
}
