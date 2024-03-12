package passage_order_table

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/global"
	"juliano.com/passage-order/interval"
	"juliano.com/passage-order/ionicon"
	mapp "juliano.com/passage-order/map"
	"juliano.com/passage-order/user"
)

var oddTdStyle = mapp.Merge(cellStyle, map[string]string{
	"background-color": "rgb(248, 247, 254)",
})

func (p *PassageOrderTable) body() []app.UI {
	TRs := []app.UI{}

	planning, _ := interval.RepeatingFromDuration(
		p.StartTime,
		p.PassageDuration,
		nil,
	)

	for index, user := range p.Group.Users {
		passageTime := *planning.NextGeneratedInterval()

		tdStyle := p.tdStyle(index)

		nameTd := app.Td().Text(user).Styles(tdStyle)
		timeTd := app.Td().Text(passageTime).Styles(tdStyle)
		actionTd := app.Td().Body(
			&ionicon.Ionicon{
				Name: "ellipsis-horizontal",
			},
		).Styles(tdStyle).Style("cursor", "pointer").
			OnClick(p.onTdActionClick(&user, index))

		tr := app.Tr().Body(
			nameTd,
			timeTd,
			actionTd,
		)

		TRs = append(TRs, tr)
	}

	return TRs
}

func (p *PassageOrderTable) tdStyle(index int) map[string]string {
	if index%2 == 1 {
		return oddTdStyle
	}

	return cellStyle
}

func (p *PassageOrderTable) onTdActionClick(user *user.User, index int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		e.StopImmediatePropagation()
		global.Popover.Show(ctx.Src(), p.tdClickItems(user, index))
	}
}
