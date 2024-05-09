package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/movie/internal/controller/movie"
	metadatagateway "github.com/ibiscum/Microservices-with-Go/Chapter0X/movie/internal/gateway/metadata/grpc"
	ratinggateway "github.com/ibiscum/Microservices-with-Go/Chapter0X/movie/internal/gateway/rating/grpc"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter0X/movie/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/pkg/discovery"
	"github.com/ibiscum/Microservices-with-Go/Chapter0X/pkg/discovery/memory"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

const serviceName = "movie"

func main() {
	f, err := os.Open("base.yaml")
	if err != nil {
		panic(err)
	}
	var cfg config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	port := cfg.API.Port
	log.Printf("Starting the movie service on port %d", port)
	registry := memory.NewRegistry()
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	defer registry.Deregister(ctx, instanceID, serviceName)
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	const limit = 100
	const burst = 100
	l := newLimiter(limit, burst)
	srv := grpc.NewServer(grpc.UnaryInterceptor(ratelimit.UnaryServerInterceptor(l)))
	reflection.Register(srv)
	gen.RegisterMovieServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

type limiter struct {
	l *rate.Limiter
}

func newLimiter(limit int, burst int) *limiter {
	return &limiter{rate.NewLimiter(rate.Limit(limit), burst)}
}

func (l *limiter) Limit() bool {
	return l.l.Allow()
}
