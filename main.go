package main

import (
	"log"
	"net/http"
	"time"
	"wednesday-app/handlers"

	"github.com/gorilla/mux"
)



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/split_tweets",handlers.BreakWords).Methods("PUT")
	r.HandleFunc("/ping",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Web server is up!`))
	}).Methods("GET")

	srv := http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
