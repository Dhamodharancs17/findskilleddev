package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func server(r *chi.Mux) {

	// Adding the middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

}
