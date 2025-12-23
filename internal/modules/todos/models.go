package todos

import "time"

type Todo struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoRepo interface {
	ChangeStatus(id string) bool
	CreateTodo()
	DeleteTodo(id string)
}
