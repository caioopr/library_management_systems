package models

import (
	"time"
)

// User represents the book data object
// TODO: Make compatible with Google Books API response OBJ.
type Book struct {
	ID               uint64    `json:"id,omitempty"`
	Title            string    `json:"name,omitempty"`
	Authors          []string  `json:"nickname,omitempty"`
	AvarageRating    float64   `json:"avarage_rating,omitempty"`
	Isbn             string    `json:"isbn,omitempty"`
	Isbn13           string    `json:"isbn13,omitempty"`
	LanguageCode     string    `json:"language_code,omitempty"`
	NumPages         uint64    `json:"num_pages,omitempty"`
	RatingsCount     uint64    `json:"ratings_count,omitempty"`
	TextReviewsCount uint64    `json:"text_reviews_count,omitempty"`
	PublicationDate  time.Time `json:"created_at,omitempty"`
	Publisher        string    `json:"publisher,omitempty"`
}
