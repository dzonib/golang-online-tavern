// Package aggregate holds our aggregates that combine many entities and VO into a full object
package aggregate

import (
	"errors"

	"github.com/google/uuid"

	"github.com/dzonib/golang-online-tavern/entity"
	"github.com/dzonib/golang-online-tavern/valueobject"
)

// ErrEmptyCustomerName - custom errors are making testing easier
var (
	ErrEmptyCustomerName = errors.New("customer name must not be empty")
)

type Customer struct {
	// person is the root entity of customer
	// that means person.ID is a main identifier for the customer

	// lowercase - aggregates should not be accessible to grab data from outside
	// no json tags also, it is not up to aggregate to decide how data should be formatted
	// we use pointers, we want data to be changed across the whole runtime when something changes
	// if we have person at multiple places, we want changes to be reflected
	person   *entity.Person
	products []*entity.Item

	// we are not using a pointer at value object because it is immutable
	transaction []valueobject.Transaction
}

// factory pattern - design patter that is used to encapsulate complex logic inside functions for creating the wanted instances
// without the caller not knowing anything about actual implementation
// DDD suggests factories for creating complex aggregates, repositories, or so.

// NewCustomer is a factory to create a new customer aggregate
// it will validate that name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrEmptyCustomerName
	}

	// root entity (we are using person id as our aggregate id)?
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	// we are creating empty values to avoid nil pointer exceptions
	// factory is helping in that way also
	return Customer{
		person:      person,
		products:    make([]*entity.Item, 0),
		transaction: make([]valueobject.Transaction, 0),
	}, nil
}

// we can not modify or get data directly from aggregate, we need to expose functions to do so

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	// we could return error here
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}
