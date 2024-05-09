package testutil

import (
	"github.com/ibiscum/Microservices-with-Go/Chapter09/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/rating/internal/controller/rating"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter09/rating/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/rating/internal/repository/memory"
)

// NewTestRatingGRPCServer creates a new rating gRPC server to be used in tests.
func NewTestRatingGRPCServer() gen.RatingServiceServer {
	r := memory.New()
	ctrl := rating.New(r, nil)
	return grpchandler.New(ctrl)
}
