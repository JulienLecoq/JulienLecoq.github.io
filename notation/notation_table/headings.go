package notation_table

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/global"
	"juliano.com/passage-order/ionicon"
	"juliano.com/passage-order/mapp"
	"juliano.com/passage-order/table"
)

func (n *NotationTable) headings() app.UI {
	style := mapp.Merge(table.ThStyle, map[string]string{
		"cursor": "pointer",
	})

	return app.Tr().Body(
		app.Th().Text("Nom").Styles(table.ThStyle),
		app.Th().Text("Note").Styles(table.ThStyle),

		app.Th().Body(ionicon.Ellipsis()).Styles(style).OnClick(n.onThActionClick),
	)
}

func (n *NotationTable) onThActionClick(ctx app.Context, e app.Event) {
	e.StopImmediatePropagation()
	global.Popover.Show(ctx.Src(), n.thClickItems())
}
