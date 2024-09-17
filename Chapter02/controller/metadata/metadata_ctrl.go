package metadata

import (
	"context"
	"errors"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/model"
	"github.com/ibiscum/Microservices-with-Go/Chapter02/repository"
)

// ErrNotFound is returned when a requested record is not found.
var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines a metadata service controller.
type MetadataController struct {
	repo metadataRepository
}

// New creates a metadata service controller.
func NewMetadataController(repo metadataRepository) *MetadataController {
	return &MetadataController{repo}
}

// Get returns movie metadata by id.
func (c *MetadataController) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
