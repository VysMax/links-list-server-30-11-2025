package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VysMax/links-list-server/entities"
)

func LinksListHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var req entities.LinksNumRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

}
