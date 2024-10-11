package middleware

import (
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins, adjust as necessary
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Check if the request is for CORS OPTIONS (pre-flight)
		if r.Method == "OPTIONS" {
			// Just add headers and send response
			w.WriteHeader(http.StatusOK)
			return
		}

		// Serve the next handler
		next.ServeHTTP(w, r)
	})
}
