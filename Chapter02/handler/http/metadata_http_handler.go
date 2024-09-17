package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/controller/metadata"
	"github.com/ibiscum/Microservices-with-Go/Chapter02/repository"
)

// MetadataHandler defines a movie metadata HTTP handler.
type MetadataHandler struct {
	ctrl *metadata.MetadataController
}

// New creates a new movie metadata HTTP handler.
func NewMetadataHandler(ctrl *metadata.MetadataController) *MetadataHandler {
	return &MetadataHandler{ctrl}
}

// GetMetadata handles GET /metadata requests.
func (h *MetadataHandler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	m, err := h.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
