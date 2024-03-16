package notation

import (
	"juliano.com/passage-order/h1"
	"juliano.com/passage-order/theme"
)

func notationSettingsHeader() *h1.H1 {
	return &h1.H1{
		Text: "Paramètres de notation",
	}
}

func notationGradeHeader() *h1.H1 {
	return &h1.H1{
		Text: "Notation",
		Styles: map[string]string{
			"margin-bottom": theme.ContentBlockMargin.String(),
		},
	}
}
