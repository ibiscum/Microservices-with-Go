package model

// Metadata defines the movie metadata.
type Metadata struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
}

// MovieDetails includes movie metadata its aggregated rating.
type MovieDetails struct {
	Rating   *float64 `json:"rating,omitempty"`
	Metadata Metadata `json:"metadata"`
}

// RecordID defines a record id. Together with RecordType identifies unique records across all types.
type RecordID string

// RecordType defines a record type. Together with RecordID identifies unique records across all types.
type RecordType string

// Existing record types.
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id.
type UserID string

// RatingValue defines a value of a rating record.
type RatingValue int

// Rating defines an individual rating created by a user for some record.
type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
