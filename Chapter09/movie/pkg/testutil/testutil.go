package testutil

import (
	"github.com/ibiscum/Microservices-with-Go/Chapter09/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/movie/internal/controller/movie"
	metadatagateway "github.com/ibiscum/Microservices-with-Go/Chapter09/movie/internal/gateway/metadata/grpc"
	ratinggateway "github.com/ibiscum/Microservices-with-Go/Chapter09/movie/internal/gateway/rating/grpc"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter09/movie/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/pkg/discovery"
)

// NewTestMovieGRPCServer creates a new movie gRPC server to be used in tests.
func NewTestMovieGRPCServer(registry discovery.Registry) gen.MovieServiceServer {
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	return grpchandler.New(ctrl)
}
