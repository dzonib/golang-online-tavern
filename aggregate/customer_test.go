package aggregate_test

import (
	"errors"
	"testing"

	"github.com/dzonib/golang-online-tavern/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test          string
		name          string
		expectedError error
	}

	testCases := []testCase{
		{test: "empty name validation", name: "", expectedError: aggregate.ErrEmptyCustomerName},
		{test: "valid name", name: "King Kong", expectedError: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedError) {
				// v stands for interface
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
