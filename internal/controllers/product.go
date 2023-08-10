package controllers

import (
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/entities/product"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/repositories/adapter"
	"github.com/google/uuid"
)

type Controller struct {
	repository adapter.Interface
}

type Interface interface {
	ListOne(ID uuid.UUID) (entity product.Product, err error)
	ListAll() (entity []product.Product, err error)
	Create(entity *product.Product) (uuid.UUID, error)
	Update(ID uuid.UUID, entity *product.Product) error
	Remove(ID uuid.UUID) error
}

func NewController(repository adapter.Interface) Interface {
	return &Controller{repository: repository}
}

func (c *Controller) ListOne(id uuid.UUID) (entity product.Product, err error) {
	entity.ID = id
	res, err := c.repository.FindOne(entity.GetFilterId(), entity.TableName())

	if err != nil {
		return entity, err
	}

	return product.ParseDynamoAttributeToStruct(res.Item)
}

func (c *Controller) ListAll() (entity []product.Product, err error) {

}

func (c *Controller) Create(entity *product.Product) (uuid.UUID, error) {

}

func (c *Controller) Update(id uuid.UUID, entity *product.Product) error {

}

func (c *Controller) Remove(id uuid.UUID) error {

}