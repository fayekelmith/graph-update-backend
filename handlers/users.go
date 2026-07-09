package handlers

import (
	"net/http"

	"github.com/fayekelmith/graph-update-backend/db"
	"github.com/go-chi/chi/v5"
)

// ListUsers replaces the former ListPeople handler (/people -> /users).
func ListUsers(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, db.AllPeople())
}

// GetUser replaces the former GetPerson handler (/people/{id} -> /users/{id}).
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, ok := db.FindPerson(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, p)
}
