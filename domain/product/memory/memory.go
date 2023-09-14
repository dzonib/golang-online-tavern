package memory

import (
	"sync"

	"github.com/dzonib/golang-online-tavern/domain/product"

	"github.com/dzonib/golang-online-tavern/aggregate"
	"github.com/google/uuid"
)

type ProductMemoryRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *ProductMemoryRepository {
	return &ProductMemoryRepository{products: make(map[uuid.UUID]aggregate.Product)}
}

func (mpr *ProductMemoryRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	// convert map into slice
	for _, p := range mpr.products {
		products = append(products, p)
	}

	return products, nil
}

// GetByID searches for a product based on it's ID
func (mpr *ProductMemoryRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (mpr *ProductMemoryRepository) Add(p aggregate.Product) error {
	// do we need this?
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	mpr.products[p.GetID()] = p

	return nil
}

func (mpr *ProductMemoryRepository) Update(p aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[p.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[p.GetID()] = p

	return nil
}

func (mpr *ProductMemoryRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mpr.products, id)

	return nil
}
