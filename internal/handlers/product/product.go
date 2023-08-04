package product

import (
	"errors"
	"net/http"

	"github.com/Asad2730/DynamoDB_CRUD_App/internal/handlers"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/handlers/product"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/repositories/adapter"
	Http "github.com/Asad2730/DynamoDB_CRUD_App/utils/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	handlers.Interface
	Controller product.Interface
	Rules      Rules.Intergace
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(repository),
		Rules:      RulesProduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	if chi.URLParam(r, "ID") != "" {
		h.GetOne(w, r)
	} else {
		h.GetAll(w, r)
	}
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {

	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		Http.StatusBadRequest(w, r, errors.New("ID is not uuid valid "))
		return
	}

	res, err := h.Controller.ListOne(ID)
	if err != nil {
		Http.StatusInternalServerError(w, r, err)
		return

	}

	Http.StatusOk(w, r, res)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	res, err := h.Controller.ListAll()
	if err != nil {
		Http.StatusInternalServerError(w, r, err)
		return

	}

	Http.StatusOk(w, r, res)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {

	body, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		Http.StatusBadRequest(w, r, err)
		return
	}

	ID, err := h.Controller.Create(body)

	if err != nil {
		Http.StatusInternalServerError(w, r, err)
		return
	}

	Http.StatusOk(w, r, map[string]interface{}{"id": ID.string()})
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {

	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		Http.StatusBadRequest(w, r, errors.New("ID is not uuid valid "))
		return
	}

	body, err := h.getBodyAndValidate(r, ID)

	if err != nil {
		Http.StatusBadRequest(w, r, err)
		return
	}

	if err := Controller.Update(ID, body); err != nil {

		Http.StatusInternalServerError(w, r, err)
		return
	}

	Http.StatusNoContent(w, r, nil)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		Http.StatusBadRequest(w, r, errors.New("ID is not uuid valid "))
		return
	}

	if err := h.Controller.Remove(ID); err != nil {
		Http.StatusInternalServerError(w, r, err)
	}

	Http.StatusNoContent(w, r, nil)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	Http.StatusNoContent(w, r, nil)
}

func (h *Handler) getBodyAndValidate(r *http.Request, ID uuid.UUID) (*EntityProduct.Product, error) {

	body := &EntityProduct.Product{}
}

func setDefaultValue() {

}
