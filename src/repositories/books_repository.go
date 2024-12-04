package repositories

import (
	"api/src/models"
	"context"
	"database/sql"
	"fmt"
)

var ctx context.Context

type booksRepository struct {
	db *sql.DB
}

func NewBooksRepository(db *sql.DB) *booksRepository {
	return &booksRepository{db}
}

func (repository booksRepository) Create(book models.Book) (uint64, error) {

	// Create a helper function for preparing failure results.
	fail := func(err error) (uint64, error) {
		return 0, fmt.Errorf("CreateOrder: %v", err)
	}

	// Get a Tx for making transaction requests.
	tx, err := repository.db.Begin()
	if err != nil {
		return fail(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// TODO: verify if the book already exists

	// inserts the authors in authors table and save their IDs into a slice
	// the id will be used with the book ID on the books_authors table
	authorsIDs := []int64{}
	for _, v := range book.Authors {
		result, err := tx.Exec("insert into authors (name) values(?)", v)
		if err != nil {
			return fail(err)
		}
		lastInsertedId, err := result.LastInsertId()
		if err != nil {
			return fail(err)
		}
		authorsIDs = append(authorsIDs, lastInsertedId)
	}

	// inserts the book in books table
	result, err := tx.Exec("insert into books (title, avarage_rating, isbn, isbn13, language_code, num_pages, ratings_count,text_reviews_count,publication_date, publisher) values (?,?,?,?,?,?,?,?,?,?)",
		book.Title,
		book.AvarageRating,
		book.Isbn,
		book.Isbn13,
		book.LanguageCode,
		book.NumPages,
		book.RatingsCount,
		book.TextReviewsCount,
		book.PublicationDate,
		book.Publisher)

	if err != nil {
		return fail(err)
	}

	bookID, err := result.LastInsertId()

	if err != nil {
		return fail(err)
	}

	// Insert the book id with each authors id into books_authors table
	for _, v := range authorsIDs {
		_, err := tx.Exec("insert into books_authors (author_id,book_id) values(?,?)", v, bookID)
		if err != nil {
			return fail(err)
		}
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	// TODO: return Book id
	return uint64(bookID), nil
}
