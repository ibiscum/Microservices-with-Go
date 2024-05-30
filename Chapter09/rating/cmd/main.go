package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ibiscum/Microservices-with-Go/Chapter09/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/pkg/discovery"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/pkg/discovery/consul"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/rating/internal/controller/rating"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter09/rating/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter09/rating/internal/repository/memory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

const serviceName = "rating"

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
	log.Printf("Starting the rating service on port %d", port)

	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer func() {
		err := registry.Deregister(ctx, instanceID, serviceName)
		if err != nil {
			log.Fatal(err)
		}
	}()
	repo := memory.New()
	ctrl := rating.New(repo, nil)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterRatingServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
