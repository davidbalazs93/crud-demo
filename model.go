package main

type Movie struct {
	Id     string `json:"movie.Id"`
	Name   string `json:"movie.Name"`
	Actors []Actor `json:"movie.Actors"`
}

type Actor struct {
	Name string `json:"actor.Name"`
}
