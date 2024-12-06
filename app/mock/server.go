package mock

import (
	"log/slog"
	"net/http"
	"time"

	"math/rand"
	"nerd3/ise_mock/mock/internaluser"
	"nerd3/ise_mock/mock/networkdevice"
)

// WithRandomDelay returns an http.HandlerFunc that introduces a random delay between 1 to 5 seconds before calling the next handler.
func WithRandomDelay(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generowanie losowego czasu opóźnienia od 1 do 5 sekund
		delay := time.Duration(rand.Intn(5)+1) * time.Second
		slog.Info("Sleeping", "seconds", delay)
		time.Sleep(delay)

		// Wywołanie oryginalnego handlera
		next(w, r)
	}
}

// WithLoggingAndHeaders is a middleware that logs the client IP and request URI, and sets the Content-Type header to application/json.
func WithLoggingAndHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.Header.Get("X-Forwarded-For")
		if clientIP == "" {
			clientIP = r.RemoteAddr
		}
		// Logowanie adresu IP
		slog.Info("Accepted connection.", "source", clientIP, "endpoint", r.RequestURI)

		// Ustawienie nagłówka Content-Type
		w.Header().Set("Content-Type", "application/json")

		// Wywołanie oryginalnego handlera
		next(w, r)
	}
}

func NewMockMux() *http.ServeMux {
	mux := http.NewServeMux()
	// internal user
	mux.HandleFunc("GET /internaluser", WithLoggingAndHeaders(WithRandomDelay(internaluser.GetInternalUserList)))
	mux.HandleFunc("GET /internaluser/{user_id}", WithLoggingAndHeaders(WithRandomDelay(internaluser.GetInternalUser)))
	// network device
	mux.HandleFunc("GET /networkdevice", WithLoggingAndHeaders(WithRandomDelay(networkdevice.GetNetworkDevice)))
	mux.HandleFunc("DELETE /networkdevice/{uuid}", WithLoggingAndHeaders(WithRandomDelay(networkdevice.DeleteNetworkDevice)))
	return mux
}
