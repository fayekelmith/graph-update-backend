package handlers

import (
	"net/http"

	"github.com/fayekelmith/graph-update-backend/db"
	"github.com/go-chi/chi/v5"
)

func GetWorkspace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	writeJSON(w, http.StatusOK, map[string]string{"id": id, "name": "Acme Workspace"})
}

// ListWorkspaceBounties returns the bounties scoped to a workspace (new in the
// evolved state; gains a frontend caller on graph-update-frontend@after).
func ListWorkspaceBounties(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "id")
	writeJSON(w, http.StatusOK, db.AllBounties())
}
