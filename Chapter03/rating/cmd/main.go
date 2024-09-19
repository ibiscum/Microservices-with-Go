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

	"github.com/ibiscum/Microservices-with-Go/Chapter03/pkg/discovery"
	"github.com/ibiscum/Microservices-with-Go/Chapter03/pkg/discovery/consul"
	"github.com/ibiscum/Microservices-with-Go/Chapter03/rating/internal/controller/rating"
	httphandler "github.com/ibiscum/Microservices-with-Go/Chapter03/rating/internal/handler/http"
	"github.com/ibiscum/Microservices-with-Go/Chapter03/rating/internal/repository/memory"
	"github.com/ibiscum/Microservices-with-Go/Chapter03/rating/pkg/model"
)

const serviceName = "rating"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "API handler port")
	flag.Parse()
	log.Printf("starting the rating service on port %d", port)
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

	log.Println("creating a new rating repository")
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

	log.Println("create a new rating controller")
	ctrl := rating.New(repo)

	log.Println("create new handler")
	h := httphandler.New(ctrl)

	log.Println("start serving")
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
