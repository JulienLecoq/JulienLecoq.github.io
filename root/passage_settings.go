package root

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func passageSettings(r *Root) app.UI {
	style := map[string]string{
		"display":        "flex",
		"flex-direction": "column",
		"gap":            "8px",
	}

	return app.Div().Body(
		r.Selector,
		r.TimePicker,
		r.PassageDuration,
	).Styles(style)
}
