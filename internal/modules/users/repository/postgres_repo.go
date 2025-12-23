package repository

import (
	"github.com/Chandra5468/todo-go/internal/modules/users"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	db *pgxpool.Pool
}

// NewRepo returns the interface defined in models, not the struct
func NewRepo(db *pgxpool.Pool) users.UserRepository {
	return &PostgresRepo{
		db: db,
	}
}
func (r *PostgresRepo) FindByID(id int) (users.User, error) {
	// r.db.Query()
	return users.User{}, nil
}
func (r *PostgresRepo) FindByUsername(username string) (users.User, error) {
	return users.User{}, nil
}
