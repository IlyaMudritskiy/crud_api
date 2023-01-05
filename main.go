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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func getMovieById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(request)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			break
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(request.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000))

	movies = append(movies, movie)

	json.NewEncoder(writer).Encode(movie)

}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			var movie Movie

			movies = append(movies[:index], movies[index+1:]...)

			_ = json.NewDecoder(request.Body).Decode(&movie)

			movie.ID = params["id"]

			movies = append(movies, movie)

			json.NewEncoder(writer).Encode(movie)
		}
	}
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(movies)
}

var movies []Movie

func main() {
	var router = mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "111111", Title: "Terminator 1", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "222222", Title: "Terminator 2", Director: &Director{Firstname: "Peter", Lastname: "Marks"}})
	movies = append(movies, Movie{ID: "3", Isbn: "333333", Title: "Terminator 3", Director: &Director{Firstname: "John", Lastname: "Galt"}})
	movies = append(movies, Movie{ID: "4", Isbn: "444444", Title: "Terminator 4", Director: &Director{Firstname: "Bernard", Lastname: "Brown"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
