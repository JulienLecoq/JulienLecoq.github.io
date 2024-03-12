package popover

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func arrow() app.UI {
	styles := map[string]string{
		"width":         "0",
		"height":        "0",
		"border-left":   "5px solid transparent",
		"border-right":  "5px solid transparent",
		"border-bottom": "5px solid black",
	}

	return app.Div().Styles(styles)
}
