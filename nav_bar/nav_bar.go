package nav_bar

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/nav_bar/routes"
)

var navBarStyle = map[string]string{
	"background-color": "#2c3e50",
	// "padding":          "16px",
}

var listStyle = map[string]string{
	"list-style":      "none",
	"padding":         "0",
	"margin":          "0",
	"display":         "flex",
	"flex-direction":  "row",
	"justify-content": "center",
	"justify-items":   "center",
	"align-content":   "center",
	"align-items":     "center",
}

var navItem = map[string]string{
	"margin": "0 12px",
}

type NavBar struct {
	app.Compo
}

func (n *NavBar) Render() app.UI {
	return app.Nav().Body(
		app.Ul().Body(
			app.Li().Body(
				app.A().Href(routes.Root).Text("Ordre de passage").Class("nav-bar-item"),
			).Styles(navItem),
			app.Li().Body(
				app.A().Href(routes.Notation).Text("Notation").Class("nav-bar-item"),
			).Styles(navItem),
		).Styles(listStyle),
	).Styles(navBarStyle)
}
