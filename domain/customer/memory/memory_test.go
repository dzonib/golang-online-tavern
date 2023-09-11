package memory

import (
	"errors"
	"github.com/dzonib/golang-online-tavern/aggregate"
	"github.com/dzonib/golang-online-tavern/domain/customer"
	"github.com/google/uuid"
	"testing"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name          string
		id            uuid.UUID
		expectedError error
	}

	// create customer
	cust, err := aggregate.NewCustomer("king")

	if err != nil {
		t.Fatal(err)
	}

	// get customers id
	id := cust.GetID()

	repo := MemoryRepository{customers: map[uuid.UUID]aggregate.Customer{
		id: cust,
	}}

	testCases := []testCase{
		{name: "no customer by id", id: uuid.MustParse("f47ac10b-58cc-4372-a567-0e02b2c3d479"), expectedError: customer.ErrCustomerNotFound},
		{name: "customer by id", id: id, expectedError: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedError) {
				// v stands for interface
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
