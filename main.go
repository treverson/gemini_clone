package main

import (
	"log"
	"net/http"

	"github.com/3cb/gemini_clone/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// serve static files
	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.PathPrefix("/dist/").Handler(http.FileServer(http.Dir("./static")))

	// trade request routes
	r.Handle("/api/trades/{product}", handlers.Trades())

	// start server
	log.Fatal(http.ListenAndServe(":4000", r))
}
