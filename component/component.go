package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/theme"
)

type Bounding struct {
	Top    theme.CssUnit
	Right  theme.CssUnit
	Bottom theme.CssUnit
	Left   theme.CssUnit
}

func BoundingClientRect(component app.UI) Bounding {
	boundingRect := component.JSValue().Call("getBoundingClientRect")

	top := boundingRect.Get("top").Float()
	bottom := boundingRect.Get("bottom").Float()
	right := boundingRect.Get("right").Float()
	left := boundingRect.Get("left").Float()

	bounding := Bounding{
		Top:    theme.CssUnit{Value: top, DefaultUnit: theme.Px},
		Right:  theme.CssUnit{Value: right, DefaultUnit: theme.Px},
		Bottom: theme.CssUnit{Value: bottom, DefaultUnit: theme.Px},
		Left:   theme.CssUnit{Value: left, DefaultUnit: theme.Px},
	}

	return bounding
}

func HtmlOf(component app.UI) string {
	return component.JSValue().Get("outerHTML").String()
}
