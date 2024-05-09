package testutil

import (
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/movie/internal/controller/movie"
	metadatagateway "github.com/ibiscum/Microservices-with-Go/Chapter0X/movie/internal/gateway/metadata/grpc"
	ratinggateway "github.com/ibiscum/Microservices-with-Go/Chapter0X/movie/internal/gateway/rating/grpc"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter0X/movie/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/pkg/discovery"
)

// NewTestMovieGRPCServer creates a new movie gRPC server to be used in tests.
func NewTestMovieGRPCServer(registry discovery.Registry) gen.MovieServiceServer {
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	return grpchandler.New(ctrl)
}
