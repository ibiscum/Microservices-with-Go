package main

import (
	"log"
	"net/http"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/controller/metadata"
	httphandler "github.com/ibiscum/Microservices-with-Go/Chapter02/handler/http"
	"github.com/ibiscum/Microservices-with-Go/Chapter02/repository/memory"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.NewMetadataRepository()
	ctrl := metadata.NewMetadataController(repo)
	h := httphandler.NewMetadataHandler(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
