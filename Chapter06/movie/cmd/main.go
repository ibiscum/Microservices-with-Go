package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ibiscum/Microservices-with-Go/Chapter06/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter06/movie/internal/controller/movie"
	metadatagateway "github.com/ibiscum/Microservices-with-Go/Chapter06/movie/internal/gateway/metadata/grpc"
	ratinggateway "github.com/ibiscum/Microservices-with-Go/Chapter06/movie/internal/gateway/rating/grpc"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter06/movie/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter06/pkg/discovery"
	"github.com/ibiscum/Microservices-with-Go/Chapter06/pkg/discovery/memory"
	"google.golang.org/grpc"
)

const serviceName = "movie"

func main() {
	var port int
	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting the movie service on port %d", port)
	registry := memory.NewRegistry()
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	defer func() {
		err := registry.Deregister(ctx, instanceID, serviceName)
		if err != nil {
			log.Panic(err)
		}
	}()
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterMovieServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
