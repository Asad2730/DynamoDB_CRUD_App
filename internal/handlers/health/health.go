package health

import (
	"errors"
	"net/http"

	"github.com/Asad2730/DynamoDB_CRUD_App/internal/handlers"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/repositories/adapter"
	Http "github.com/Asad2730/DynamoDB_CRUD_App/utils/http"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	if !h.Repository.Health() {
		Http.StatusInternalServerError(w, r, errors.New("relational Database not alive"))
		return
	}

	Http.StatusOk(w, r, "Service Ok")

}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {

	Http.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {

	Http.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	Http.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {

	Http.StatusNoContent(w, r, nil)
}
