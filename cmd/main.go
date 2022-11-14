package main

import (
	"log"
	"net/http"

	"github.com/briares/hello-api/handlers"
	"github.com/briares/hello-api/handlers/rest"
)

func main() {

	addr := ":8080" //

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler) //
	mux.HandleFunc("/health", handlers.HealthCheck) //

	log.Printf("listening on %s\n", addr) //

	log.Fatal(http.ListenAndServe(addr, mux)) //
}
