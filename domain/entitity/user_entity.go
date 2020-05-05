package entitity

import (
	"time"
)

// User entity
type User struct {
	ID uint64
	EmailAddress string
	FirstName string
	LastName string
	Password string
	RegisteredAt time.Time
	ModifiedAt time.Time
}