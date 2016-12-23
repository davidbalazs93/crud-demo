package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func handleDefaultPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func displayAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from [%q] to retrieve the list of movies", r.RemoteAddr)
	movieService := neo4jMovieService{}
	moviesJSON, err := json.Marshal(movieService.getMovies())
	if (err != nil) {
		log.Fatalf("Error while converting the array of movies to JSON: [%v]", err)
	}
	log.Printf("%s", moviesJSON)
	w.Write(moviesJSON)
}

func getById(w http.ResponseWriter, r *http.Request) {
	if (len(r.URL.Path) < 8) {
		log.Print("The url path is too short. It does not contain the movie id.")
		w.Write([]byte("Provide a movie id."))
		return
	}

	movieId := r.URL.Path[8:]
	log.Printf("Received request from [%q] to retrieve movie with id [%s]", r.RemoteAddr, movieId)
	movieService := neo4jMovieService{}
	movie, err := movieService.getMovieById(movieId)
	if (err != nil) {
		log.Printf("Cannot retrieve movie with id: %s, because [%v]", movieId, err)
		return
	}

	movieJSON, err := json.Marshal(movie)

	if (err != nil) {
		log.Fatalf("Cannot marshal movie [%v]", err)
	}

	log.Printf("Retrieving movie: %s", movieJSON)
	w.Write(movieJSON)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	log.Print("Received request to add a new movie.")
	decoder := json.NewDecoder(r.Body)
	var newMovie Movie
	err := decoder.Decode(&newMovie)
	if (err != nil) {
		log.Fatalf("Cannot unmarshal request body [%v]", err)
	}
	movieService := neo4jMovieService{}
	movieService.add(newMovie)
}
