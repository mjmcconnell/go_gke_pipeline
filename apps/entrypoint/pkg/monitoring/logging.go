package monitoring

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{ResponseWriter: w}
}

func (rw *loggingResponseWriter) Status() int {
	return rw.status
}

func (rw *loggingResponseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.New()
		logger.Formatter = new(logrus.JSONFormatter)

		logger.WithFields(logrus.Fields{
			"url":            r.URL.String(),
			"method":         r.Method,
			"referer":        r.Referer(),
			"user_agent":     r.UserAgent(),
			"content_length": r.ContentLength,
		}).Info("Request")

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)
		logger.WithFields(logrus.Fields{
			"status":   wrapped.status,
			"duration": time.Since(start),
		}).Info("Response")
	})
}
