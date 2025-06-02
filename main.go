package main

import (
	"log"
	"net/http"
)

// A Server defines parameters for running an HTTP server.
// type Server struct {
// 	// Addr optionally specifies the TCP address for the server to listen on, in the form "host:port".
// 	Addr string
// }

func main() {
	// Create a new servemux
	mux := http.NewServeMux()

	// Register handler functions for specific paths
	// mux routes incoming HTTP requests to the correct handler function based on request path
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// todo
	})

	// Create a new http.Server struct (note, this can be confusing if you are use to creating your own structs as `Server` is already created for you)
	s := &http.Server{
		Addr:    ":8080", // Set the .Addr field to ":8080"
		Handler: mux,     // Use the new "ServeMux" as the server's handler
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// bootdev run 861ada77-c583-42c8-a265-657f2c453103
