package customer

import (
	"errors"

	"github.com/ayadnez/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("customer was not found")
	ErrFailedToADDCustomer    = errors.New("failed to add new customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
	Delete(uuid.UUID) error
}
