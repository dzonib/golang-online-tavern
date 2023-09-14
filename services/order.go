package services

import (
	log "
	"context"
	"github.com/dzonib/golang-online-tavern/domain/customer/mongo"

	""

	"github.com/dzonib/golang-online-tavern/domain/customer/memory"
	"github.com/dzonib/golang-online-tavern/domain/product"
	prodmemory "github.com/dzonib/golang-online-tavern/domain/product/memory"
	"github.com/google/uuid"

	"github.com/dzonib/golang-online-tavern/aggregate"

	"github.com/dzonib/golang-online-tavern/domain/customer"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

// we can specifyrepositories we want to use like this
//NewOrderService(WithCustomerRepository, WithMemoryProductRepository)

// factory function for our service
//
// variable ammount of order configurations

// NewOrderService takes a variable number of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the order-service
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches order configuration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings, they could be inputted here
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connectionString)

		if err != nil {
			return nil
		}

		os.customers = cr
		return nil
	}
}

// WithMemoryProductRepository adds an in memory product repo and adds all input products
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := prodmemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

//func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
//	return func(os *OrderService) error {
//		pr := prodmemory.New()
//
//		for _, p := range products {
//			if err := pr.Add(p); err != nil {
//				return err
//			}
//		}
//
//		return nil
//	}
//}

// CreateOrder will chain-together all repositories to create an order for a customer
// will return the collected price of all Products
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product
	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exist in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	// Add Products and Update Customer

	return price, nil
}
