package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"aevyn/finops-arch/internal/application"
)

func main() {
	app := application.New()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      app.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println(time.Now(), "server init on 8080; glhf")
	log.Fatal(server.ListenAndServe())
}
