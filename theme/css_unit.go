package theme

import (
	"fmt"
)

type CssUnit struct {
	Value       float64
	DefaultUnit Unit
}

func (u *CssUnit) String() string {
	return fmt.Sprintf("%v%s", u.Value, u.DefaultUnit)
}

func (u *CssUnit) ToPx() string {
	return fmt.Sprintf("%v%s", u.Value, "px")
}

func (u *CssUnit) ToRem() string {
	return fmt.Sprintf("%v%s", u.Value, "rem")
}

func (u *CssUnit) Add(value float64, opts ...OperatorOption) CssUnit {
	newUnit := CssUnit{u.Value + value, u.DefaultUnit}

	for _, opt := range opts {
		opt(&newUnit)
	}

	return newUnit
}

func (u *CssUnit) AddUnit(value CssUnit, opts ...OperatorOption) CssUnit {
	return u.Add(value.Value, opts...)
}

func (u *CssUnit) Mult(value float64) CssUnit {
	return CssUnit{u.Value * value, u.DefaultUnit}
}

// Option defines a function that applies a configuration to CssUnit
type OperatorOption func(*CssUnit)

// WithUnit is an option to set the unit of CssUnit
func WithUnit(u Unit) OperatorOption {
	return func(c *CssUnit) {
		c.DefaultUnit = u
	}
}
