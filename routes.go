package main

import (
	"github.com/fayekelmith/graph-update-backend/handlers"
	"github.com/go-chi/chi/v5"
)

// NewRouter wires every HTTP endpoint of the bounty API.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Bounties
	r.Get("/bounties", handlers.ListBounties)
	r.Post("/bounties", handlers.CreateBounty)
	r.Get("/bounties/{id}", handlers.GetBounty)
	r.Put("/bounties/{id}", handlers.UpdateBounty)
	r.Delete("/bounties/{id}", handlers.DeleteBounty)

	// People
	r.Get("/people", handlers.ListPeople)
	r.Get("/people/{id}", handlers.GetPerson)

	// Auth
	r.Post("/auth/login", handlers.Login)

	// Workspaces — intentionally has no frontend caller (dangling endpoint).
	r.Get("/workspaces/{id}", handlers.GetWorkspace)

	return r
}
