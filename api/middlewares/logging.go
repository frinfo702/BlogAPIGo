package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

// constructor
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// override WriteHeader methond
func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

// func myMiddleware(next http.Handler) http.Handler
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// loging request information
		log.Println(req.URL, req.Method) // print the request URL and method

		rlw := NewResLoggingWriter(w)
		// return the response
		next.ServeHTTP(rlw, req) // pass the request to the next handler
		// logging the content of responses
		log.Println("res: ", rlw.code)
	})
}
