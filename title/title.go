package title

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/theme"
)

type Title struct {
	app.Compo

	Text string
}

func (t *Title) Render() app.UI {
	return app.Div().Text(t.Text).
		Styles(
			map[string]string{
				"width":  "fit-content",
				"margin": fmt.Sprintf("0 auto %s auto", &theme.ContentBlockMargin),
				// "color":     "rgb(181, 199, 92)",
				"color":     "#003366",
				"font-size": "3.8rem",
			},
		)
}
