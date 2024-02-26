package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Response struct {
	Films []Film `json:"films"`
}

type Film struct {
	ReleaseYear int    `json:"release_year"`
	Title       string `json:"title"`
	Producer    string `json:"producer"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func FilmsHandler(w http.ResponseWriter, r *http.Request) {
	// declare response variable
	var response Response

	// Retrieve film details
	films := prepareResponse()

	// assign film details to response
	response.Films = films

	// update content type
	w.Header().Set("Content-Type", "application/json")

	// specify HTTP status code
	w.WriteHeader(http.StatusOK)

	// convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update response
	w.Write(jsonResponse)
}

func prepareResponse() []Film {
	var films []Film

	film := Film{ReleaseYear: 2001, Title: "The Lord of the Rings", Producer: "Peter Robert Jackson"}
	films = append(films, film)

	film = Film{ReleaseYear: 2011, Title: "Game of Thrones", Producer: "Neil Marshall"}
	films = append(films, film)

	film = Film{ReleaseYear: 1999, Title: "The Green Mile", Producer: "Frank Darabont"}
	films = append(films, film)

	film = Film{ReleaseYear: 1978, Title: "Midnight Express", Producer: "Alan Parker"}
	films = append(films, film)

	film = Film{ReleaseYear: 2006, Title: "The Prestige", Producer: "Charles R. Rogers"}
	films = append(films, film)

	return films
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/films", FilmsHandler).Methods("GET")
	http.Handle("/", router)

	// start and listen to requests
	http.ListenAndServe(":8080", router)
}