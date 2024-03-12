package h1

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/theme"
)

type H1 struct {
	app.Compo

	Text   string
	Styles map[string]string
}

func (h *H1) Render() app.UI {
	return app.H1().
		Text(h.Text).
		Styles(
			map[string]string{
				"margin": fmt.Sprintf("%s 0 %s 0", &theme.ContentBlockMargin, &theme.ContentBlockMargin),
				// "color":     "rgb(105, 177, 179)",
				"color":     "#0055A4",
				"font-size": "1.8rem",
			},
		).
		Styles(h.Styles)
}
