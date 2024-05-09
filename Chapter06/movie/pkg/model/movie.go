package model

import "github.com/ibiscum/Microservices-with-Go/Chapter06/metadata/pkg/model"

// MovieDetails includes movie metadata its aggregated rating.
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
