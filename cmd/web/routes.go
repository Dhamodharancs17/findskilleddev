package main

import (
	"net/http"

	"github.com/Dhamodharancs17/findskilleddev/internal/organization"
	"github.com/Dhamodharancs17/findskilleddev/internal/user"
	"github.com/go-chi/chi/v5"
)

func router() {
	r := chi.NewRouter()

	// Adding the middleware
	server(r)

	// Defining the routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	//User routes
	r.Post("/user/create", user.CreateUser)
	r.Get("/user/show", user.GetUser)

	// Organization routes
	r.Post("/organization/create", organization.CreateOrganization)
	r.Get("/organization/show", organization.GetOrganization)

	// Start the server
	http.ListenAndServe(":3000", r)
}
