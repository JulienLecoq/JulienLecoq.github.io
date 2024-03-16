package passage_order_table

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/global"
	"juliano.com/passage-order/interval"
	"juliano.com/passage-order/ionicon"
	"juliano.com/passage-order/table"
	"juliano.com/passage-order/user"
)

func (p *PassageOrderTable) body() []app.UI {
	TRs := []app.UI{}

	planning, _ := interval.RepeatingFromDuration(
		p.StartTime,
		p.PassageDuration,
		nil,
	)

	for index, user := range p.Group.Users {
		passageTime := *planning.NextGeneratedInterval()

		tdStyle := table.TdStyle(index)

		nameTd := app.Td().Text(user).Styles(tdStyle)
		timeTd := app.Td().Text(passageTime).Styles(tdStyle)
		actionTd := app.Td().Body(ionicon.Ellipsis()).Styles(table.IconStyleFromStyle(tdStyle)).
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

func (p *PassageOrderTable) onTdActionClick(user *user.User, index int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		e.StopImmediatePropagation()
		global.Popover.Show(ctx.Src(), p.tdClickItems(user, index))
	}
}
