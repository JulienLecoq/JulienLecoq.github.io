package button

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Button struct {
	app.Compo

	Text string
}

func (b *Button) SetText(text string) {
	b.Text = text
	b.Update()
}

func (b *Button) Render() app.UI {
	style := map[string]string{
		"border-radius":    "10px",
		"background-color": "blue",
	}

	return app.
		Button().
		Text(b.Text).
		Styles(style)
}
