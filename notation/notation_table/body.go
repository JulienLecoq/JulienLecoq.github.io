package notation_table

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/global"
	"juliano.com/passage-order/ionicon"
	"juliano.com/passage-order/notation/graded_user"
	"juliano.com/passage-order/table"
)

func (n *NotationTable) body() []app.UI {
	TRs := []app.UI{}

	for index, user := range n.Group.Users {
		tdStyle := table.TdStyle(index)

		nameTd := app.Td().Text(user).Styles(tdStyle)
		noteTd := app.Td().Body(n.noteInput(user)).Styles(tdStyle)

		actionTd := app.Td().Body(ionicon.Ellipsis()).Styles(table.IconStyleFromStyle(tdStyle)).
			OnClick(n.onTdActionClick(user, index))

		tr := app.Tr().Body(
			nameTd,
			noteTd,
			actionTd,
		)

		TRs = append(TRs, tr)
	}

	return TRs
}

func (n *NotationTable) noteInput(user *graded_user.GradedUser) app.UI {
	style := map[string]string{
		"width": "6rem",
		"padding": "5px",
	}

	return app.
		Input().
		Type("number").
		Value(user.Grade).
		Min(0).
		Max(20).
		Required(true).
		OnInput(func(ctx app.Context, e app.Event) {
			value := ctx.JSSrc().Get("value").String()
			grade, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return
			}

			prevGrade := user.Grade
			user.Grade = grade

			if n.OnChange != nil {
				n.OnChange(user, prevGrade, user.Grade)
			}
		}).
		Styles(style)
}

func (n *NotationTable) onTdActionClick(user *graded_user.GradedUser, index int) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		e.StopImmediatePropagation()
		global.Popover.Show(ctx.Src(), n.tdClickItems(user, index))
	}
}
