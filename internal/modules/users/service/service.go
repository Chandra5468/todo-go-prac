package service

import (
	"github.com/Chandra5468/todo-go/internal/modules/users"
	"github.com/jackc/pgx/v5/pgxpool"
)

// If you want u can separate repository and service logic if needed

type Service interface {
	GetUser(id int) (users.User, error)
}

type svc struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) Service {
	return &svc{
		db: db,
	}
}

func (s *svc) GetUser(id int) (users.User, error) {
	return users.User{}, nil
}
