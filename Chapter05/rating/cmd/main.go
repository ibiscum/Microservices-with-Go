package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ibiscum/Microservices-with-Go/Chapter05/gen"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/pkg/discovery"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/pkg/discovery/consul"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/rating/internal/controller/rating"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter05/rating/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/rating/internal/repository/memory"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/rating/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "rating"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.Parse()
	log.Printf("starting the rating service on port %d", port)
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
				log.Println("failed to report healthy state: " + err.Error())
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

	log.Println("initialize rating repository with some values")
	var entries []*model.Rating

	fileData, err := os.ReadFile("./rating_entries.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileData, &entries)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		err := repo.Put(context.Background(), model.RecordID(entry.RecordID), model.RecordType(entry.RecordType), entry)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctrl := rating.New(repo)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterRatingServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
