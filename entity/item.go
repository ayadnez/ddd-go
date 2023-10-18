package entity

import (
	"github.com/google/uuid"
)

type Itemt struct {
	ID          uuid.UUID
	Name        string
	Description string
}
