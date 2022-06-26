package queries

import (
	"github.com/google/uuid"
	"github.com/hotrungnhan/go-fiber-template/app/models"
	"gorm.io/gorm"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*gorm.DB
}

// GetUserByID query for getting one User by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE id = $1`

	// Send query to database.
	tx := q.Raw(query, id).Scan(user)
	if tx.Error != nil {
		// Return empty object and error.
		return user, tx.Error
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	// Define User variable.
	user := models.User{}

	// Define query string.
	query := `SELECT * FROM users WHERE email = $1`

	// Send query to database.
	tx := q.Raw(query, email).Scan(user)
	if tx.Error != nil {
		// Return empty object and error.
		return user, tx.Error
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Define query string.
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Send query to database.
	tx := q.Exec(
		query,
		u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.PasswordHash, u.UserStatus, u.UserRole,
	)
	if tx.Error != nil {
		// Return only error.
		return tx.Error
	}

	// This query returns nothing.
	return nil
}
