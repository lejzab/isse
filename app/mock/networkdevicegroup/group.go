package networkdevicegroup

import (
	"math/rand"
	"net/http"
)

// CreateGroup handles an HTTP request and responds with a created status or an internal server error at a random probability.
func CreateGroup(w http.ResponseWriter, r *http.Request) {

	if rand.Float32() < 0.25 {
		w.WriteHeader(http.StatusInternalServerError) // group already exists
		return
	}
	w.WriteHeader(http.StatusCreated)
}
