package main

import (
	"log"
	"net/http"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/movie/internal/controller/movie"
	metadatagateway "github.com/ibiscum/Microservices-with-Go/Chapter02/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/ibiscum/Microservices-with-Go/Chapter02/movie/internal/gateway/rating/http"
	httphandler "github.com/ibiscum/Microservices-with-Go/Chapter02/movie/internal/handler/http"
)

func main() {
	log.Println("starting the movie service")

	log.Println("create the gateways")
	metadataGateway := metadatagateway.New("http://localhost:8081")
	ratingGateway := ratinggateway.New("http://localhost:8082")

	log.Println("create a new movie controller")
	ctrl := movie.New(ratingGateway, metadataGateway)

	log.Println("create new handler")
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))

	log.Println("start serving")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
