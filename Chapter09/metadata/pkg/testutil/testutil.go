package testutil

import (
	"github.com/ibiscum/Microservices-with-Go/Chapter09/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/metadata/internal/controller/metadata"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter09/metadata/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/metadata/internal/repository/memory"
)

// NewTestMetadataGRPCServer creates a new metadata gRPC server to be used in tests.
func NewTestMetadataGRPCServer() gen.MetadataServiceServer {
	r := memory.New()
	ctrl := metadata.New(r)
	return grpchandler.New(ctrl)
}
