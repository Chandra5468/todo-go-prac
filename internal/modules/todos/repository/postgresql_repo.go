package repository

import (
	"github.com/Chandra5468/todo-go/internal/modules/todos"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreRepo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) todos.TodoRepo {
	return &PostgreRepo{
		db: db,
	}
}

func (p *PostgreRepo) ChangeStatus(id string) bool {
	return false
}
func (p *PostgreRepo) CreateTodo() {

}
func (p *PostgreRepo) DeleteTodo(id string) {

}
