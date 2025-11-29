package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VysMax/links-list-server/entities"
)

func (h *LinksHandler) LinksList(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var req entities.LinksNumRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	h.service.GetLinks(req)

	w.Header().Set("content-type", "application/pdf")
	// _, err = w.Write([]byte(sb.String()))
	if err != nil {
		fmt.Println(err)
	}
}
