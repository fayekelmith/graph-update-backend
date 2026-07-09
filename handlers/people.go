package handlers

import (
	"net/http"

	"github.com/fayekelmith/graph-update-backend/db"
	"github.com/go-chi/chi/v5"
)

func ListPeople(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, db.AllPeople())
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, ok := db.FindPerson(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, p)
}
