package movie

import (
	"context"
	"errors"
	"sync"

	metadatamodel "github.com/ibiscum/Microservices-with-Go/Chapter13/metadata/pkg/model"
	"github.com/ibiscum/Microservices-with-Go/Chapter13/movie/internal/gateway"
	"github.com/ibiscum/Microservices-with-Go/Chapter13/movie/pkg/model"
	ratingmodel "github.com/ibiscum/Microservices-with-Go/Chapter13/rating/pkg/model"
)

// ErrNotFound is returned when the movie metadata is not found.
var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

// Controller defines a movie service controller.
type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New creates a new movie service controller.
func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway, metadataGateway}
}

// Get returns the movie details including the aggregated rating and movie metadata.
func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var metadata *metadatamodel.Metadata
	var getMetadataErr error
	var rating float64
	var getRatingErr error
	go func() {
		defer wg.Done()
		metadata, getMetadataErr = c.metadataGateway.Get(ctx, id)
	}()
	go func() {
		defer wg.Done()
		rating, getRatingErr = c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)
	}()
	wg.Wait()
	if err := getMetadataErr; err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	details := &model.MovieDetails{Metadata: *metadata}

	if err := getRatingErr; err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Just proceed in this case, it's ok not to have ratings yet.
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
