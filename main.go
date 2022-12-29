package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{Firstname: "Yash", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "448744", Title: "Movie Two", Director: &Director{Firstname: "John", Lastname: "Deol"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
