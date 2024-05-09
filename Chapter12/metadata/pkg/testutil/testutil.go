package testutil

import (
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/metadata/internal/controller/metadata"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter0X/metadata/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/metadata/internal/repository/memory"
)

// NewTestMetadataGRPCServer creates a new metadata gRPC server to be used in tests.
func NewTestMetadataGRPCServer() gen.MetadataServiceServer {
	r := memory.New()
	ctrl := metadata.New(r)
	return grpchandler.New(ctrl)
}
