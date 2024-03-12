package signal

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Signal[T any] struct {
	component app.Compo
	value     T
}

func New[T any](value T, comp app.Compo) *Signal[T] {
	return &Signal[T]{
		value:     value,
		component: comp,
	}
}

func (s *Signal[T]) get() T {
	return s.value
}

func (s *Signal[T]) set(value T) {
	s.value = value
	s.component.Update()
}
