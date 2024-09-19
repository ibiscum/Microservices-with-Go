package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/metadata/internal/controller/metadata"
	httphandler "github.com/ibiscum/Microservices-with-Go/Chapter02/metadata/internal/handler/http"
	memory "github.com/ibiscum/Microservices-with-Go/Chapter02/metadata/internal/repository/memory"
	model "github.com/ibiscum/Microservices-with-Go/Chapter02/metadata/pkg/model"
)

func main() {
	log.Println("starting the movie metadata service")

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
		err := repo.Put(context.Background(), entry.ID, entry)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("create a new metadata controller")
	ctrl := metadata.New(repo)

	log.Println("create new handler")
	h := httphandler.New(ctrl)
	http.Handle("/metadata/{id}", http.HandlerFunc(h.GetMetadata))

	log.Println("start serving")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
