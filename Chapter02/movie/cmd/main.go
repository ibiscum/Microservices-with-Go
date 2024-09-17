package main

import (
	"log"
	"net/http"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/controller/movie"
	metadatagateway "github.com/ibiscum/Microservices-with-Go/Chapter02/gateway/metadata/http"
	ratinggateway "github.com/ibiscum/Microservices-with-Go/Chapter02/gateway/rating/http"
	httphandler "github.com/ibiscum/Microservices-with-Go/Chapter02/handler/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := movie.NewMovieController(ratingGateway, metadataGateway)
	h := httphandler.NewMovieHandler(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
