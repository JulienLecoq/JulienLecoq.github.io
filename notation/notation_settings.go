package notation

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func notationSettings(r *NotationPage) app.UI {
	style := map[string]string{
		"display":        "flex",
		"flex-direction": "column",
		"gap":            "8px",
	}

	return app.Div().Body(
		r.Selector,
	).Styles(style)
}
