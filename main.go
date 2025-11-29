package main

import (
	"log"
	"net/http"

	"github.com/VysMax/links-list-server/entities"
	controller "github.com/VysMax/links-list-server/handlers"
	"github.com/VysMax/links-list-server/repository"
	"github.com/VysMax/links-list-server/usecase"
)

func main() {

	entities.SavedLinksDir = "./saved_links"

	repo := repository.New()
	usecase := usecase.New(repo)
	handler := controller.New(usecase)

	http.HandleFunc("/links", handler.SaveLinks)
	http.HandleFunc("/links_list", handler.LinksList)
	http.HandleFunc("/pdf_file", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=kittens.pdf")
		w.Header().Set("content-type", "application/pdf")
		w.Write([]byte("trfd;gls;dflg;sldfkg'waf"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error starting server:%v", err)
	}
}
