package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)

	lookupEnv, envPortExists := os.LookupEnv("PORT")

	var err error

	if envPortExists {
		log.Printf("Server launched using port %s", lookupEnv)
		err = http.ListenAndServe(":"+lookupEnv, nil)
	} else {
		log.Println("Server launched using port 8081")
		err = http.ListenAndServe(":8081", nil)
	}

	if err != nil {
		log.Fatalf("Error %s", err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user")
	if userName == "" {
		userName = "guest"
	}
	_, err := fmt.Fprintf(w, "Hello, %s!", userName)

	if err != nil {
		return
	}
}
