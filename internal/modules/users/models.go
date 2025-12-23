package users

import "time"

// 1. The Data Structure
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// 2. The Interface (The Port)
// This defines what the repository MUST do.
type UserRepository interface {
	FindByID(id int) (User, error)
	FindByUsername(username string) (User, error)
}
