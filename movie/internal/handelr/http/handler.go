package http

import (
	"encoding/json"
	"errors"
	"log"
	"movieapp/movie/internal/controller/movie"
	"net/http"
)

type Handler struct {
	ctrl *movie.Controller
}

func New(ctrl *movie.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetMovieDetails(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	details, err := h.ctrl.Get(req.Context(), id)
	if err != nil {
		if errors.Is(err, movie.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Responce encode error: %v\n", err)
	}
}
