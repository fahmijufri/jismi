package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/fahmijufri/jismi/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func HTTPStatLogger() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()

		next(w, r)

		if !IsHealthcheckURL(r.URL.String()) {
			responseTime := time.Now()
			deltaTime := responseTime.Sub(startTime).Seconds() * 1000

			logrus.WithFields(logrus.Fields{
				"request_time":  startTime.Format(time.RFC3339),
				"delta_time":    deltaTime,
				"response_time": responseTime.Format(time.RFC3339),
				"request_proxy": r.RemoteAddr,
				"url":           r.URL.Path,
				"method":        r.Method,
			}).Infoln("Request")
		}
	}
}

func IsHealthcheckURL(url string) bool {
	switch {
	case strings.Contains(url, config.PathLiveness):
		fallthrough
	case strings.Contains(url, config.PathReadiness):
		return true
	default:
		return false
	}
}
