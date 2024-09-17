package main

import (
	"log"
	"net/http"

	"github.com/ibiscum/Microservices-with-Go/Chapter02/controller/rating"
	httphandler "github.com/ibiscum/Microservices-with-Go/Chapter02/handler/http"
	"github.com/ibiscum/Microservices-with-Go/Chapter02/repository/memory"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.NewRatingRepository()
	ctrl := rating.NewRatingController(repo)
	h := httphandler.NewRatingHandler(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
