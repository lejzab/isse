package internaluser

// Definicje struktur odpowiadających danym JSON
type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
	Type string `json:"type"`
}

type Resource struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        Link   `json:"link"`
}

type SearchResult struct {
	Total     int        `json:"total"`
	Resources []Resource `json:"resources"`
}

type InternalUserList struct {
	SearchResult SearchResult `json:"SearchResult"`
}
