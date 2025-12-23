package delivery

import (
	"net/http"

	"github.com/Chandra5468/todo-go/internal/modules/users/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc service.Service
}

func NewHandler(r chi.Router, svc service.Service) {
	h := &Handler{
		svc: svc,
	}

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", h.GetUser)
	})
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

}
