package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func router(){
	r := chi.NewRouter()

	// Adding the middleware
	server(r)

	// Defining the routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	//User routes
	

	
	// Start the server
	http.ListenAndServe(":3000", r)
}