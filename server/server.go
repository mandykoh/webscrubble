package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/go-chi/chi"
	"github.com/mandykoh/webscrubble/api"
	"log"
	"os"
	"strconv"
)

func main() {
	const FrontendRoot = "frontend"

	port := 8080
	if len(os.Args) > 1 {
		parsedPort, err := strconv.ParseInt(os.Args[1], 10, 16)
		if err != nil {
			log.Fatalf("Unable to parse port '%s'", os.Args[1])
			os.Exit(1)
		}

		port = int(parsedPort)
	}

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

	log.Printf("Starting server on port %d", port)

	log.Fatal(server.ListenAndServe())
}
