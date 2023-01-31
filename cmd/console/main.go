package main

import (
	"fmt"
	api "github.com/Lefree111/gorilla-mux/rest-api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.Home).Methods("GET")
	r.HandleFunc("/api/v1/albums", api.PostAlbum).Methods("POST")
	r.HandleFunc("/api/v1/albums", api.GetAlbums).Methods("GET")
	r.HandleFunc("/api/v1/albums/{id}", api.GetAlbum).Methods("GET")
	r.HandleFunc("/api/v1/albums/{id}", api.UpdateAlbum).Methods("PUT")
	r.HandleFunc("/api/v1/albums/{id}", api.DeleteAlbum).Methods("DELETE")
	fmt.Println("Listening and serving on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
