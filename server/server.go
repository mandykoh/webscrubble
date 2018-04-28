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

	const DefaultHTTPPort = 8080
	const DefaultHTTPSPort = 8443
	const TLSCertFile = "cert.pem"
	const TLSPrivateKeyFile = "privkey.pem"

	useTLS := true
	if _, err := os.Stat(TLSCertFile); err != nil {
		useTLS = false
		log.Printf("Couldn't find %s", TLSCertFile)
	}
	if _, err := os.Stat(TLSPrivateKeyFile); err != nil {
		useTLS = false
		log.Printf("Couldn't find %s", TLSPrivateKeyFile)
	}

	port := DefaultHTTPPort
	if useTLS {
		port = DefaultHTTPSPort
	}

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

	if useTLS {
		log.Printf("Starting HTTPS server on port %d", port)
		log.Fatal(server.ListenAndServeTLS(TLSCertFile, TLSPrivateKeyFile))
	} else {
		log.Printf("Starting HTTP server on port %d", port)
		log.Fatal(server.ListenAndServe())
	}
}
