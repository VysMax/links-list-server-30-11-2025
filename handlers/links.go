package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VysMax/links-list-server/entities"
)

func LinksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var req entities.LinksRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	var resp entities.LinksResponse

	checkedLinks, num, err := checkAvailability(req.Links)
	if err != nil {
		http.Error(w, "failed to make HTTP request to check availability", http.StatusInternalServerError)
	}

	resp.Links = checkedLinks
	resp.LinksNum = num

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
	}
}
