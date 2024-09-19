package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ibiscum/Microservices-with-Go/Chapter03/metadata/internal/controller/metadata"
	httphandler "github.com/ibiscum/Microservices-with-Go/Chapter03/metadata/internal/handler/http"
	"github.com/ibiscum/Microservices-with-Go/Chapter03/metadata/internal/repository/memory"
	"github.com/ibiscum/Microservices-with-Go/Chapter03/metadata/pkg/model"
	"github.com/ibiscum/Microservices-with-Go/Chapter03/pkg/discovery"
	"github.com/ibiscum/Microservices-with-Go/Chapter03/pkg/discovery/consul"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()

	log.Printf("starting the metadata service on port %d\n", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
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

	log.Println("creating a new metadata repository")
	repo := memory.New()

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

	log.Println("create a new metadata controller")
	ctrl := metadata.New(repo)

	log.Println("create new handler")
	h := httphandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))

	log.Println("start serving")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
