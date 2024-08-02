package Utils

import (
	"github.com/fatih/color"
	"log"
	"net/http"
	"time"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	ErrorMsg   string
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// LoggingMiddleware logs the incoming HTTP requests with color
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		crw := &CustomResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
		methodColor := color.New(color.FgGreen).SprintFunc()
		urlColor := color.New(color.FgCyan).SprintFunc()
		next.ServeHTTP(crw, r)
		durationColor := color.New(color.FgYellow).SprintFunc()
		statusColor := color.New(color.FgRed).SprintFunc()
		if crw.StatusCode < 400 {
			statusColor = color.New(color.FgGreen).SprintFunc()
		}
		if crw.StatusCode >= 400 {
			log.Printf("%s %s %s %s in %s - Error: %s", methodColor(r.Method), urlColor(r.RequestURI), "Status-code", statusColor(crw.StatusCode), durationColor(time.Since(start)), crw.ErrorMsg)
		} else {
			log.Printf("%s %s %s %s in %s", methodColor(r.Method), urlColor(r.RequestURI), "Status-code", statusColor(crw.StatusCode), durationColor(time.Since(start)))
		}
	})
}
