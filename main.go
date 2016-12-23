package main

import (
	"github.com/jawher/mow.cli"
	"log"
	"strconv"
	"os"
	"net/http"
)

func main() {
	cliApp := cli.App("Crud app", "CRUD app to test go skills")
	port := cliApp.Int(cli.IntOpt{
		Name: "port",
		Value: 8080,
		Desc: "Port app. 8080 by default",
		EnvVar: "PORT",
	})

	cliApp.Action = func() {
		log.Printf("Listening on port: %d", *port)
		listen(*port)
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatalf("Cannot start app [%v]", err)
	}
}

func listen(port int) {
	http.HandleFunc("/", handleDefaultPage)
	http.HandleFunc("/display-all", displayAll)
	http.HandleFunc("/movies/", getById)
	http.HandleFunc("/add-movie", addMovie)
	http.HandleFunc("/actors/", displayActorsForMovie)

	err := http.ListenAndServe(":" + strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
