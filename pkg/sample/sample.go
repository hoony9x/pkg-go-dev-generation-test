package sample

import (
	"math/rand"
)

// Sample defines sample 2 methods.
type Sample interface {
	// GetName returns existing name.
	GetName() string

	// SetName replaces existing name.
	SetName(name string)
}

// sample contains 2 fields
type sample struct {
	// name holds string value
	name string

	// randomNumber holds int value
	randomNumber int
}

// GetName returns existing name.
func (s *sample) GetName() string {
	return s.name
}

// SetName replaces existing name.
func (s *sample) SetName(name string) {
	s.name = name
}

// NewSample creates new sample object with given and random number.
func NewSample(name string) Sample {
	return &sample{
		name:         name,
		randomNumber: rand.Intn(100),
	}
}
