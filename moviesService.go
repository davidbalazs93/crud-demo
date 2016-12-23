package main

type moviesService interface {
	getMovies() []Movie
	betMovieById(id string) (Movie, error)
	add(movie Movie)
}
