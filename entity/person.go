package entity

import (
	"github.com/google/uuid"
)

type Person struct {
	ID      uuid.UUID
	Age     int
	Name    string
	address string
}
