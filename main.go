package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies = []Movie{
	{Id: "1", Isbn: "4890213", Title: "Movie one", Director: &Director{Firstname: "John", Lastname: "Doe"}},
	{Id: "2", Isbn: "4892421", Title: "Movie two", Director: &Director{Firstname: "Davin", Lastname: "Larsen"}},
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-type")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-type")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-type")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-type")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movie.Director.Firstname = "Sam"
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-type")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {

		}
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/:id", getMovie).Methods("GET")
	r.HandleFunc("/movies/:id", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/:id", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
