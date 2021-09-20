package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dariobanfi/svg-flag-generator-with-name/getflag"
)

func main() {
	http.HandleFunc("/", getflag.GetFlag)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
