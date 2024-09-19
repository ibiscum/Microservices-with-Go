package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/rating/internal/controller/rating"
	"github.com/ibiscum/Microservices-with-Go/Chapter02/rating/pkg/model"
)

// Handler defines a HTTP rating handler.
type Handler struct {
	ctrl *rating.Controller
}

// New creates a new rating service HTTP handler.
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl}
}

// Handle handles PUT and GET /rating requests.
func (h *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	pType := req.PathValue("type")
	recordType := model.RecordType(pType)
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pId := req.PathValue("id")
	recordID := model.RecordID(pId)
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(req.Context(), recordID, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
	case http.MethodPut:
		userID := model.UserID(req.FormValue("userId"))
		v, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := h.ctrl.PutRating(req.Context(), recordID, recordType, &model.Rating{UserID: userID, Value: model.RatingValue(v)}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
