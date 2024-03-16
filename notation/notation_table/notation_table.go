package notation_table

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/mapp"
	"juliano.com/passage-order/notation/graded_user"
	"juliano.com/passage-order/table"
)

type NotationTable struct {
	app.Compo

	Group *graded_user.GradedGroup

	Styles   map[string]string
	OnChange func(user *graded_user.GradedUser, prevGrade float64, newGrade float64)
}

func (n *NotationTable) Render() app.UI {
	TRs := append([]app.UI{n.headings()}, n.body()...)
	style := mapp.Merge(table.TableStyle, n.Styles)

	return app.Table().Body(TRs...).Styles(style)
}
