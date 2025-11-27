package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/VysMax/links-list-server/entities"
)

func LinksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if entities.LastNum == 0 {
		entities.LastNum = entities.ExtractLastNum(entities.SavedLinksDir)
	}

	var req entities.LinksRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	var linksSet entities.LinksResponse

	checkedLinks, err := checkAvailability(req.Links)
	if err != nil {
		http.Error(w, "failed to make HTTP request to check availability", http.StatusInternalServerError)
	}

	linksSet.Links = checkedLinks
	entities.LastNum++
	linksSet.LinksNum = entities.LastNum

	resp, err := json.Marshal(linksSet)
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
	}

	err = os.WriteFile("./saved_links/set_"+strconv.Itoa(entities.LastNum)+".json", resp, 0644)
	if err != nil {
		http.Error(w, "failed to create file with saved links set", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(resp); err != nil {
		http.Error(w, "failed to write saved links set", http.StatusInternalServerError)
	}
}
