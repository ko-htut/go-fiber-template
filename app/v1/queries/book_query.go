package queries

import (
	"github.com/google/uuid"
	"github.com/hotrungnhan/go-fiber-template/app/models"
	"gorm.io/gorm"
)

// BookQueries struct for queries from Book model.
type BookQueries struct {
	*gorm.DB
}

// GetBooks method for getting all books.
func (q *BookQueries) GetBooks() ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	query := `SELECT * FROM books`

	// Send query to database.
	tx := q.Raw(query).Scan(&books)

	if tx.Error != nil {
		// Return empty object and error.
		return books, tx.Error
	}

	// Return query result.
	return books, nil
}

// GetBooksByAuthor method for getting all books by given author.
func (q *BookQueries) GetBooksByAuthor(author string) ([]models.Book, error) {
	// Define books variable.
	books := []models.Book{}

	// Define query string.
	query := `SELECT * FROM books WHERE author = $1`

	// Send query to database.
	tx := q.Raw(query, author).Scan(&books)

	if tx.Error != nil {
		// Return empty object and error.
		return books, tx.Error
	}

	// Return query result.
	return books, nil
}

// GetBook method for getting one book by given ID.
func (q *BookQueries) GetBook(id uuid.UUID) (models.Book, error) {
	// Define book variable.
	book := models.Book{}

	// Define query string.
	query := `SELECT * FROM books WHERE id = $1`

	// Send query to database.
	tx := q.Raw(query, id).Scan(&book)

	if tx.Error != nil {
		// Return empty object and error.
		return book, tx.Error
	}

	// Return query result.
	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookQueries) CreateBook(b *models.Book) error {
	// Define query string.
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	tx := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Title, b.Author, b.BookStatus, b.BookAttrs)

	if tx.Error != nil {
		// Return empty object and error.
		return tx.Error
	}
	// This query returns nothing.
	return nil
}

// UpdateBook method for updating book by given Book object.
func (q *BookQueries) UpdateBook(id uuid.UUID, b *models.Book) error {
	// Define query string.
	query := `UPDATE books SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

	// Send query to database.
	tx := q.Exec(query, id, b.UpdatedAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if tx.Error != nil {
		// Return only error.
		return tx.Error
	}

	// This query returns nothing.
	return nil
}

// DeleteBook method for delete book by given ID.
func (q *BookQueries) DeleteBook(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	tx := q.Exec(query, id)
	if tx.Error != nil {
		// Return only error.
		return tx.Error
	}

	// This query returns nothing.
	return nil
}
