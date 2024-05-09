package model

import "github.com/ibiscum/Microservices-with-Go/Chapter0X/metadata/pkg/model"

// MovieDetails includes movie metadata its aggregated rating.
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata`
}
