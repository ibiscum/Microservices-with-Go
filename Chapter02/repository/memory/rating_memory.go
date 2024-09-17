package memory

import (
	"context"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/model"
	"github.com/ibiscum/Microservices-with-Go/Chapter02/repository"
)

// Repository defines a rating repository.
type RatingRepository struct {
	data map[model.RecordType]map[model.RecordID][]model.Rating
}

// New creates a new memory repository.
func NewRatingRepository() *RatingRepository {
	return &RatingRepository{map[model.RecordType]map[model.RecordID][]model.Rating{}}
}

// Get retrieves all ratings for a given record.
func (r *RatingRepository) Get(ctx context.Context, recordID model.RecordID, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, repository.ErrNotFound
	}
	if ratings, ok := r.data[recordType][recordID]; !ok || len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}
	return r.data[recordType][recordID], nil
}

// Put adds a rating for a given record.
func (r *RatingRepository) Put(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = map[model.RecordID][]model.Rating{}
	}
	r.data[recordType][recordID] = append(r.data[recordType][recordID], *rating)
	return nil
}
