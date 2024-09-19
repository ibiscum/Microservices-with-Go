package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/rating/internal/controller/rating"
	httphandler "github.com/ibiscum/Microservices-with-Go/Chapter02/rating/internal/handler/http"
	memory "github.com/ibiscum/Microservices-with-Go/Chapter02/rating/internal/repository/memory"
	model "github.com/ibiscum/Microservices-with-Go/Chapter02/rating/pkg/model"
)

func main() {
	log.Println("starting the rating service")

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
	http.Handle("/rating", http.HandlerFunc(h.Handle))

	log.Println("start serving")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
