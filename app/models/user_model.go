package models

// User struct to describe User object.
type User struct {
	Model
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash,omitempty"`
	UserStatus   int    `json:"user_status"`
	UserRole     string `json:"user_role"`
}
