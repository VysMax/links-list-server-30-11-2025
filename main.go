package main

import (
	"log"
	"net/http"

	"github.com/VysMax/links-list-server/entities"
	"github.com/VysMax/links-list-server/handlers"
)

func main() {

	entities.SavedLinksDir = "./saved_links"

	http.HandleFunc("/links", handlers.LinksHandler)
	http.HandleFunc("/links_list", handlers.LinksListHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error starting server:%v", err)
	}
}
