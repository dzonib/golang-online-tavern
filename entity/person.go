// Package entity holds all entities that are shared across subdomains
package entity

import "github.com/google/uuid"

// entities are mutable

// Person is an entity that represents a person in all subdomains
type Person struct {
	// ID is an identifier of the entity
	// it is nice to create unique identifiers
	ID   uuid.UUID
	Name string
	Age  int
}
