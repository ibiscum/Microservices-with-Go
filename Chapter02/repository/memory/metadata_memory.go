package memory

import (
	"context"
	"sync"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/model"
	"github.com/ibiscum/Microservices-with-Go/Chapter02/repository"
)

// MetadataRepository defines a memory movie matadata repository.
type MetadataRepository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// NewMetadataRepository creates a new memory repository.
func NewMetadataRepository() *MetadataRepository {
	return &MetadataRepository{data: map[string]*model.Metadata{}}
}

// Get retrieves movie metadata by movie id.
func (r *MetadataRepository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

// Put adds movie metadata for a given movie id.
func (r *MetadataRepository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
