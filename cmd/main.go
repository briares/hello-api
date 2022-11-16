package main

import (
	"log"
	"net/http"
	"time"

	"github.com/briares/hello-api/handlers"
	"github.com/briares/hello-api/handlers/rest"
	"github.com/briares/hello-api/translation"
)

func main() {

	addr := ":8080" //

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/hello", translateHandler.TranslateHandler) //
	mux.HandleFunc("/health", handlers.HealthCheck)             //

	log.Printf("listening on %s\n", addr) //

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}
	log.Fatal(server.ListenAndServe()) //
}
