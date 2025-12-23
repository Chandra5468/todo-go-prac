package service

import (
	"github.com/Chandra5468/todo-go/internal/modules/todos"
)

type Service interface {
	GetTodoList()
}

type svc struct {
	repo todos.TodoRepo
}

func NewService(repo todos.TodoRepo) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) GetTodoList() {

}
