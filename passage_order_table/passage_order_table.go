package passage_order_table

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/theme"
	"juliano.com/passage-order/user"
)

var cellStyle = map[string]string{
	"padding":          "17px 42px 17px 42px",
	"text-align":       "left",
	"color":            "#303030",
	"background-color": "white",
	"font-size":        "0.92rem",
}

var tableStyle = map[string]string{
	"border-collapse": "collapse",
	// "box-shadow":      "rgba(100, 100, 111, 0.2) 0px 7px 29px 0px",
	"box-shadow":    "rgba(99, 99, 99, 0.2) 0px 2px 8px 0px",
	"border-radius": theme.SmallRadius.String(),
	"overflow":      "hidden",
}

type PassageOrderTable struct {
	app.Compo

	Group           user.Group
	StartTime       time.Time
	PassageDuration time.Duration

	Styles map[string]string
}

func (p *PassageOrderTable) Render() app.UI {
	TRs := append([]app.UI{p.headings()}, p.body()...)

	return app.Table().Body(
		TRs...,
	).Styles(tableStyle).Styles(p.Styles)
}
