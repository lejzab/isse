package mock

import (
	"log/slog"
	"nerd3/ise_mock/mock/networkdevicegroup"
	"net/http"
	"os"
	"strconv"
	"time"

	"math/rand"
	"nerd3/ise_mock/mock/internaluser"
	"nerd3/ise_mock/mock/networkdevice"
)

// getDelayFromEnv retrieves the "ISE_DELAY" environment variable, parses it as an integer, and returns it if valid and positive.
// If the variable is not set or invalid, it defaults to a delay of 4 seconds.
func getDelayFromEnv() int {
	// Pobierz zmienną środowiskową ISE_DELAY
	envDelay := os.Getenv("ISE_DELAY")
	if envDelay != "" {
		if parsedDelay, err := strconv.Atoi(envDelay); err == nil && parsedDelay > 0 {
			return parsedDelay
		}
	}
	// Domyślna wartość: 4 sekundy
	return 4
}

// WithRandomDelay returns an http.HandlerFunc that introduces a random delay between 1 to 5 seconds before calling the next handler.
func WithRandomDelay(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generowanie losowego czasu opóźnienia od 1 do 5 sekund
		baseDelay := getDelayFromEnv()
		delay := time.Duration(rand.Intn(baseDelay)+1) * time.Second
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

// NewMockMux creates and returns a new HTTP ServeMux with predefined routes and middleware for internal user and network device endpoints.
func NewMockMux() *http.ServeMux {
	mux := http.NewServeMux()
	// internal user
	mux.HandleFunc("GET /internaluser", WithLoggingAndHeaders(WithRandomDelay(internaluser.GetInternalUserList)))
	mux.HandleFunc("GET /internaluser/{user_id}", WithLoggingAndHeaders(WithRandomDelay(internaluser.GetInternalUser)))
	// network device
	mux.HandleFunc("GET /networkdevice", WithLoggingAndHeaders(WithRandomDelay(networkdevice.GetNetworkDeviceList)))
	mux.HandleFunc("GET /networkdevice/{uuid}", WithLoggingAndHeaders(WithRandomDelay(networkdevice.GetNetworkDevice)))
	mux.HandleFunc("PUT /networkdevice/{uuid}", WithLoggingAndHeaders(WithRandomDelay(networkdevice.UpdateNetworkDevice)))
	mux.HandleFunc("DELETE /networkdevice/{uuid}", WithLoggingAndHeaders(WithRandomDelay(networkdevice.DeleteNetworkDevice)))
	mux.HandleFunc("POST /networkdevice", WithLoggingAndHeaders(WithRandomDelay(networkdevice.CreateDevice)))
	// network device grouo
	mux.HandleFunc("POST /networkdevicegroup", WithLoggingAndHeaders(WithRandomDelay(networkdevicegroup.CreateGroup)))
	return mux
}
