package delivery

import (
	"github.com/Chandra5468/todo-go/internal/modules/todos/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc service.Service
}

func NewHandler(r chi.Router, svc service.Service) {

}
