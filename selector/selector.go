package selector

import (
	"fmt"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Selector[T fmt.Stringer] struct {
	app.Compo

	SelectorId    string
	SelectName    string
	Label         string
	Elements      []T
	SelectedIndex int
	OnChange      func(value T)
}

func (s *Selector[T]) Render() app.UI {
	return app.
		Div().
		Body(
			app.Label().For(s.SelectorId).Text(s.Label),

			app.Select().
				OnChange(s.onValueSelected).
				ID(s.SelectorId).
				Name(s.SelectName).
				Text(s.Label).Body(
				s.options(),
			),
		)
}

func (s *Selector[T]) SelectedElement() T {
	return s.Elements[s.SelectedIndex]
}

func (s *Selector[T]) options() app.RangeLoop {
	return app.Range(s.Elements).Slice(func(index int) app.UI {
		value := s.Elements[index]
		isSelected := index == s.SelectedIndex

		return app.Option().Value(index).Selected(isSelected).Text(value)
	})
}

func (s *Selector[T]) onValueSelected(ctx app.Context, e app.Event) {
	indexStr := ctx.JSSrc().Get("value").String()

	// Should never return an error since the value is set as an index to an array
	index, _ := strconv.ParseUint(indexStr, 10, 64)

	value := s.Elements[index]

	if s.OnChange != nil {
		s.OnChange(value)
	}
}
