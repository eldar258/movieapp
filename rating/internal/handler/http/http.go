package http

import (
	"encoding/json"
	"errors"
	"log"
	"movieapp/rating/internal/controller/rating"
	"movieapp/rating/pkg/model"
	"net/http"
	"strconv"
)

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	recordId := model.RecordID(req.FormValue("id"))
	if recordId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recordType := model.RecordType(req.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		value, err := h.ctrl.GetAggregatedRating(req.Context(), recordId, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(value); err != nil {
			log.Printf("Responce encode error: %v\n", err)
		}
	case http.MethodPost:
		userId := model.UserID(req.FormValue("userId"))
		value, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.ctrl.PutRating(req.Context(), recordId, recordType, &model.Rating{UserID: userId, RatingValue: model.RatingValue(value)}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
