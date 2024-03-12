package theme

// import (
// 	"fmt"
// 	"golang.org/x/exp/constraints"
// )
//
// type Number interface {
// 	constraints.Integer | constraints.Float
// }
//
// type CssUnit[T Number] struct {
// 	Value       T
// 	DefaultUnit Unit
// }
//
// func (u *CssUnit[T]) String() string {
// 	return fmt.Sprintf("%d%s", u.Value, u.DefaultUnit)
// }
//
// func (u *CssUnit[T]) toPx() string {
// 	return fmt.Sprintf("%d%s", u.Value, "px")
// }
//
// func (u *CssUnit[T]) toRem() string {
// 	return fmt.Sprintf("%d%s", u.Value, "rem")
// }
//
// func (u *CssUnit[T]) add(value T, opts ...OperatorOption[T]) CssUnit[T] {
// 	newUnit := CssUnit[T]{u.Value + value, u.DefaultUnit}
//
// 	for _, opt := range opts {
// 		opt(&newUnit)
// 	}
//
// 	return newUnit
// }
//
// func (u *CssUnit[T]) addUnit(value CssUnit[T], opts ...OperatorOption[T]) CssUnit[T] {
// 	return u.add(value.Value, opts...)
// }
//
// func (u *CssUnit[T]) mult(value T) CssUnit[T] {
// 	return CssUnit[T]{u.Value * value, u.DefaultUnit}
// }
//
// // Option defines a function that applies a configuration to CssUnit
// type OperatorOption[T Number] func(*CssUnit[T])
//
// // WithUnit is an option to set the unit of CssUnit
// func WithUnit[T Number](u Unit) OperatorOption[T] {
// 	return func(c *CssUnit[T]) {
// 		c.DefaultUnit = u
// 	}
// }
