package main

import (
	"github.com/jmcvetta/neoism"
	"errors"
	"log"
)

type neo4jMovieService struct{}

func (neo4jMovieService *neo4jMovieService) getMovies() []Movie {
	db, err := neoism.Connect("http://neo4j:david@localhost:7474/db/data")
	if (err != nil) {
		log.Fatalf("Cannot connect to DB: [%v]", err)
	}

	var movies []Movie
	cypherQuery := neoism.CypherQuery{
		Statement: `MATCH(movie:Movie) RETURN movie.Id, movie.Name`,
		Result: &movies,
	}

	db.Cypher(&cypherQuery)
	return movies
}

func (neo4jMovieService *neo4jMovieService) getMovieById(id string) (Movie, error) {
	db, err := neoism.Connect("http://neo4j:david@localhost:7474/db/data")
	if (err != nil) {
		log.Fatalf("Cannot connect to DB: [%v]", err)
	}

	var movies []Movie
	cypherQuery := neoism.CypherQuery{
		Statement: `MATCH(movie:Movie) WHERE movie.Id={id} RETURN movie.Id, movie.Name`,
		Parameters:neoism.Props{"id":id},
		Result: &movies,
	}

	db.Cypher(&cypherQuery)

	if len(movies) == 0 {
		return Movie{}, errors.New("Cannot find movie with id: " + id)
	}

	return movies[0], nil
}

func (neo4jMovieService *neo4jMovieService) add(movie Movie) {
}

func (neo4jMovieService *neo4jMovieService) getActorsForMovie(id string) ([]Actor, error) {
	db, err := neoism.Connect("http://neo4j:david@localhost:7474/db/data")
	if (err != nil) {
		log.Fatalf("Cannot connect to DB: [%v]", err)
	}

	var actors []Actor
	cypherQuery := neoism.CypherQuery{
		Statement:`match(movie:Movie {Id:{id}}) match((actor:Actor)-[r:ACTS_IN]->(movie)) return actor.Name`,
		Parameters:neoism.Props{"id":id},
		Result: &actors,
	}

	db.Cypher(&cypherQuery)

	return actors, nil
}
