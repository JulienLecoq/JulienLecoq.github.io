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
	return app.
		Button().
		Body(
			app.Text(b.Text),
		).
		Style("border-radius", "10px").
		Style("background-color", "blue")
}
