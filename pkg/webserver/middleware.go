package webserver

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sunraylab/verbose"
)

// Logger is a middleware handler that does request logging
// https://drstearns.github.io/tutorials/gomiddleware/
type Logger struct {
	handler  http.Handler
	apicalls int
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	isapi := strings.HasPrefix(r.URL.Path, "/api")
	if isapi {
		l.apicalls++
		call := l.apicalls
		verbose.Printf(verbose.TRACK, "%s %s \033[7m[#%v[\033[0m\n", r.Method, r.RequestURI, l.apicalls)
		l.handler.ServeHTTP(w, r)
		verbose.Track(start, "%s %s \033[7m]#%v]\033[0m", r.Method, r.RequestURI, call)
	} else {
		verbose.Printf(verbose.INFO, "%s %s ", r.Method, r.RequestURI)
		l.handler.ServeHTTP(w, r)
		fmt.Println("✓")
	}
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap, 0}
}

// middlewareJsonContent set "Content-Type", "application/json; charset=utf-8" header for all request starting by /api
func middlewareJsonContent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
		}
		next.ServeHTTP(w, r)
	})
}

// middlewareNoCache set cache header
func middlewareNoCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Remove cache in the answer
		w.Header().Add("Cache-Control", "no-cache")
		// Call the next handler (another middleware in the chain or the final handler)
		next.ServeHTTP(w, r)
	})
}
