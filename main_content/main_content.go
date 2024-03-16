package main_content

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/theme"
)

var style = map[string]string{
	"padding": theme.ContentMainMargin,
}

type MainContent struct {
	app.Compo

	content []app.UI
}

func NewMain() *MainContent {
	return &MainContent{}
}

func (m *MainContent) Body(content ...app.UI) app.UI {
	m.content = content

	return app.Div().Body(content...).Styles(style)
}
