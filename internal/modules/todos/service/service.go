package service

import "github.com/jackc/pgx/v5/pgxpool"

type Service interface {
	GetTodoList()
}

type svc struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) Service {
	return &svc{
		db: db,
	}
}

func (s *svc) GetTodoList() {

}
