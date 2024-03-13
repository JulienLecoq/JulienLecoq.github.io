package passage_order_table

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/global"
	"juliano.com/passage-order/ionicon"
	"juliano.com/passage-order/mapp"
)

var thStyle = mapp.Merge(cellStyle, map[string]string{
	// "background-color": "#f2f2f2",
	// "background-color": "#e0f2fe", // Previous value
	// "background-color": "#bae6fd",
	"background-color": "rgb(112, 123, 211)",
	"color":            "white",
	"font-size":        "1rem",
})

func (p *PassageOrderTable) headings() app.UI {
	return app.Tr().Body(
		app.Th().Text("Nom").Styles(thStyle),
		app.Th().Text("Heure de passage").Styles(thStyle),

		app.Th().Body(
			&ionicon.Ionicon{
				Name: "ellipsis-horizontal",
			},
		).Styles(thStyle).Style("cursor", "pointer").
			OnClick(p.onThActionClick),
	)
}

func (p *PassageOrderTable) onThActionClick(ctx app.Context, e app.Event) {
	e.StopImmediatePropagation()
	global.Popover.Show(ctx.Src(), p.thClickItems())
}
