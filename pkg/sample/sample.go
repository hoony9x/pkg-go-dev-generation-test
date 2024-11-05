package sample

import (
	"math/rand"
)

type Sample interface {
	GetName() string
	SetName(name string)
}

type sample struct {
	name         string
	randomNumber int
}

func (s *sample) GetName() string {
	return s.name
}

func (s *sample) SetName(name string) {
	s.name = name
}

func NewSample(name string) Sample {
	return &sample{
		name:         name,
		randomNumber: rand.Intn(100),
	}
}
