package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fayekelmith/graph-update-backend/db"
	"github.com/fayekelmith/graph-update-backend/models"
	"github.com/go-chi/chi/v5"
)

func ListBounties(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, db.AllBounties())
}

func CreateBounty(w http.ResponseWriter, r *http.Request) {
	var b models.Bounty
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.SaveBounty(b)
	writeJSON(w, http.StatusCreated, b)
}

func GetBounty(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	b, ok := db.FindBounty(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, b)
}

func UpdateBounty(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var b models.Bounty
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.ID = id
	db.SaveBounty(b)
	writeJSON(w, http.StatusOK, b)
}

// AssignBounty assigns a bounty to a person (new in the evolved state).
func AssignBounty(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body struct {
		PersonID string `json:"personId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, ok := db.FindBounty(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	b.Status = "assigned:" + body.PersonID
	db.SaveBounty(b)
	writeJSON(w, http.StatusOK, b)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
