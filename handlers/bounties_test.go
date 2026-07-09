package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListBounties(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/bounties", nil)
	w := httptest.NewRecorder()

	ListBounties(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestLoginRejectsBadBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/auth/login", nil)
	w := httptest.NewRecorder()

	Login(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 on empty body, got %d", w.Code)
	}
}
