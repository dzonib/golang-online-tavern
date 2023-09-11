package product

import (
	"errors"

	"github.com/dzonib/golang-online-tavern/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound = errors.New("product with that ID does not exist")
)

// manage and handle product aggregate

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetById(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
