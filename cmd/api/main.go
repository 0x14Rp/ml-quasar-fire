package main

import (
	"log"
	"ml-quasar-fire/internal/infrastructure/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/topsecret/", server.TopSecretHandler)
	mux.HandleFunc("/topsecret_split/", server.TopSecretSplitHandler)
	mux.HandleFunc("/clear_secrets/", server.ClearSatelliteStore)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
