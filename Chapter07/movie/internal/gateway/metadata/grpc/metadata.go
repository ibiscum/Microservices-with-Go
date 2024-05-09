package grpc

import (
	"context"

	"github.com/ibiscum/Microservices-with-Go/Chapter07/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter07/internal/grpcutil"
	"github.com/ibiscum/Microservices-with-Go/Chapter07/metadata/pkg/model"
	"github.com/ibiscum/Microservices-with-Go/Chapter07/pkg/discovery"
)

// Gateway defines a movie metadata gRPC gateway.
type Gateway struct {
	registry discovery.Registry
}

// New creates a new gRPC gateway for a movie metadata service.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

// Get returns movie metadata by a movie id.
func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metadata", g.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{MovieId: id})
	if err != nil {
		return nil, err
	}
	return model.MetadataFromProto(resp.Metadata), nil
}
