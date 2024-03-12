package root

import (
	"juliano.com/passage-order/h1"
	"juliano.com/passage-order/theme"
)

func passageSettingsHeader() *h1.H1 {
	return &h1.H1{
		Text: "Paramètres de passage",
	}
}

func passageOrderHeader() *h1.H1 {
	return &h1.H1{
		Text: "Ordre de passage",
		Styles: map[string]string{
			"margin-bottom": theme.ContentBlockMargin.String(),
		},
	}
}
