package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/go-chi/chi"
	"github.com/mandykoh/webscrubble/api"
	"log"
)

func main() {
	const FrontendRoot = "frontend"

	port := 8080

	endpoints := api.Endpoints{}
	fileServer := http.FileServer(http.Dir(FrontendRoot))

	router := chi.NewRouter()
	router.Get("/api/version", endpoints.Version)
	router.Get("/*", fileServer.ServeHTTP)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
