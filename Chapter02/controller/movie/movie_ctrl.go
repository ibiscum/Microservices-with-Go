package movie

import (
	"context"
	"errors"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/gateway"
	model "github.com/ibiscum/Microservices-with-Go/Chapter02/model"
)

// ErrNotFound is returned when the movie metadata is not found.
var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines a movie service controller.
type MovieController struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New creates a new movie service controller.
func NewMovieController(ratingGateway ratingGateway, metadataGateway metadataGateway) *MovieController {
	return &MovieController{ratingGateway, metadataGateway}
}

// Get returns the movie details including the aggregated rating and movie metadata.
func (c *MovieController) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, model.RecordID(id), model.RecordTypeMovie)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Just proceed in this case, it's ok not to have ratings yet.
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
