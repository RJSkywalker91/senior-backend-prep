package transport

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"playerapi/internal/player"
)

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type PlayerHandler struct{ svc *player.Service }

func NewPlayerHandler(svc *player.Service) *PlayerHandler {
	return &PlayerHandler{svc: svc}
}

func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	var req RegisterRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	params := player.RegisterParams{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	id, err := h.svc.CreatePlayer(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(fmt.Appendf(nil, "user created with id: %s", id))
}

func (h *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	p, err := h.svc.GetPlayer(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(p)
}
