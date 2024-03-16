package notation

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/theme"
)

func GradeSummary(n *NotationPage) app.UI {
	style := map[string]string{
		"display":        "flex",
		"flex-direction": "column",
		"gap":            "8px",
		"margin-bottom":  theme.ContentBlockMargin.String(),
	}

	return app.Div().Body(
		n.Average,
		n.Median,
	).Styles(style)
}
