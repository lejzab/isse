package internaluser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"log/slog"
)

func GetInternalUserList(w http.ResponseWriter, r *http.Request) {
	// Struktura do której zmapujemy JSON
	var userList InternalUserList

	// Deserializacja JSON do struktury
	err := json.Unmarshal([]byte(InternalUserListJSON), &userList)
	if err != nil {
		slog.Warn("Błąd podczas parsowania JSON", "error", err)

	}
	responseJSON, _ := json.Marshal(userList)
	w.Write([]byte(responseJSON))
}

func GetInternalUser(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf(InternalUserJSON, r.PathValue("user_id"))
	w.Write([]byte(response))
}

func GetFilteredInternalUserList(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")
	if !strings.HasPrefix(filter, "name.EQ") {
		http.Error(w, "Invalid filter", http.StatusBadRequest)
		return
	}
	name := strings.TrimPrefix(filter, "name.EQ.")

}
