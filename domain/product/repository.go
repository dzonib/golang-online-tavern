package product

import (
	"errors"

	"github.com/dzonib/golang-online-tavern/aggregate"
	"github.com/google/uuid"
)

var (
	//ErrProductNotFound is returned when a product is not found
	ErrProductNotFound = errors.New("the product was not found")
	//ErrProductAlreadyExists is returned when trying to add a product that already exists
	ErrProductAlreadyExists = errors.New("the product already exists")
)

// ProductRepository is the repository interface to fulfill to use the product aggregate
type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
