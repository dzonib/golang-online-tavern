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
	for _, product := range mpr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mpr *ProductMemoryRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
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
