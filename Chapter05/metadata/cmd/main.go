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
	"github.com/ibiscum/Microservices-with-Go/Chapter05/metadata/internal/controller/metadata"
	grpchandler "github.com/ibiscum/Microservices-with-Go/Chapter05/metadata/internal/handler/grpc"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/metadata/internal/repository/memory"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/metadata/pkg/model"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/pkg/discovery"
	"github.com/ibiscum/Microservices-with-Go/Chapter05/pkg/discovery/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("starting the metadata service on port %d\n", port)

	log.Println("create a consul service registry")
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	log.Printf("create instance of service and register: %s\n", serviceName)
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		log.Fatal(err)
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

	log.Println("create a metadata repository in memory")
	repo := memory.New()

	log.Println("create a metadata controller")
	ctrl := metadata.New(repo)

	log.Println("initialize metadata repository with some values")
	var entries []*model.Metadata

	fileData, err := os.ReadFile("./metadata_entries.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileData, &entries)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		err := repo.Put(ctx, entry.ID, entry)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("create a gRPC handler")
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("create a gRPC server")
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterMetadataServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
