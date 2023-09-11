package aggregate

import (
	"errors"
	"github.com/dzonib/golang-online-tavern/entity"
	"github.com/google/uuid"
)

var (
	ErrMissingValue = errors.New("missing name or description")
)

type Product struct {
	// root entity
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}

	return Product{item: &entity.Item{ID: uuid.New(), Name: name, Description: description}, price: price, quantity: 0}, nil
}

// what we need to expose:

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

// extract entity from aggregate
func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
