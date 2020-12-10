package monitoring

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

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
		}).Info("Incoming request")

		next.ServeHTTP(w, r)
	})
}
