package models

import (
	"errors"
	"time"
)

// User represents the book data object
// TODO: Make compatible with Google Books API response OBJ.
type Book struct {
	ID               uint64    `json:"id,omitempty"`
	Title            string    `json:"title,omitempty"`
	Authors          []string  `json:"authors,omitempty"`
	AvarageRating    float64   `json:"avarage_rating,omitempty"`
	Isbn             string    `json:"isbn,omitempty"`
	Isbn13           string    `json:"isbn13,omitempty"`
	LanguageCode     string    `json:"language_code,omitempty"`
	NumPages         uint64    `json:"num_pages,omitempty"`
	RatingsCount     uint64    `json:"ratings_count,omitempty"`
	TextReviewsCount uint64    `json:"text_reviews_count,omitempty"`
	PublicationDate  time.Time `json:"publication_date,omitempty"`
	Publisher        string    `json:"publisher,omitempty"`
}

// TODO: Make a custom "Not empy field" error
// TODO: Date correct validation
func (book *Book) Validate() error {
	if book.Title == "" {
		return errors.New("The title field can not be empty.")
	}

	if len(book.Authors) == 0 {
		return errors.New("The authors field can not be empty.")
	}
	if book.Isbn == "" {
		return errors.New("The isbn field can not be empty.")
	}

	if book.Isbn13 == "" {
		return errors.New("The isbn-13 field can not be empty.")
	}

	if book.LanguageCode == "" {
		return errors.New("The language code field can not be empty.")
	}

	if book.PublicationDate.String() == "" {
		return errors.New("The publication field can not be empty.")
	}

	if book.Publisher == "" {
		return errors.New("The publisher field can not be empty.")
	}

	return nil
}
