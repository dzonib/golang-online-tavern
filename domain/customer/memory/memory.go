// Packakge memory is in-memory implementation of Customer repository
package memory

import (
	"fmt"
	"github.com/dzonib/golang-online-tavern/domain/customer"
	"sync"

	"github.com/dzonib/golang-online-tavern/aggregate"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	// protect this with sync.Mutex
	sync.Mutex
}

// New is a factory for MemoryRepository
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if c, ok := mr.customers[id]; ok {
		return c, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {

	// if a map is not initialized,
	// factory function should have protected us against an empty customers map,
	// but for extra safety we will create it if it does not exist
	if mr.customers == nil {
		mr.Lock()
		defer mr.Unlock()

		mr.customers = make(map[uuid.UUID]aggregate.Customer)
	}

	// make sure the customer is not already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	defer mr.Unlock()

	mr.customers[c.GetID()] = c

	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	mr.Lock()
	defer mr.Unlock()

	mr.customers[c.GetID()] = c

	return nil
}
