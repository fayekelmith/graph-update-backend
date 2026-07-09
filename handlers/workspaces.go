package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetWorkspace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	writeJSON(w, http.StatusOK, map[string]string{"id": id, "name": "Acme Workspace"})
}
