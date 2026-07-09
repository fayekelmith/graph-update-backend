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
	r.Post("/bounties/{id}/assign", handlers.AssignBounty) // new

	// Users (renamed from /people)
	r.Get("/users", handlers.ListUsers)
	r.Get("/users/{id}", handlers.GetUser)

	// Auth
	r.Post("/auth/login", handlers.Login)

	// Workspaces — GetWorkspace stays dangling; workspace-bounties is new.
	r.Get("/workspaces/{id}", handlers.GetWorkspace)
	r.Get("/workspaces/{id}/bounties", handlers.ListWorkspaceBounties) // new

	return r
}
