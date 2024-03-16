package notation

import (
	// "fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/global"
	"juliano.com/passage-order/main_content"
	"juliano.com/passage-order/nav_bar"
	"juliano.com/passage-order/notation/graded_user"
	"juliano.com/passage-order/notation/notation_table"
	"juliano.com/passage-order/popover"
	"juliano.com/passage-order/selector"
	"juliano.com/passage-order/title"
)

type NotationPage struct {
	app.Compo

	Popover       *popover.Popover
	Title         *title.Title
	Selector      *selector.Selector[*graded_user.GradedGroup]
	NotationTable *notation_table.NotationTable
	Average       *Average
	Median        *Median
}

func (r *NotationPage) OnInit() {
	// r.Popover = global.Popover
	r.Popover = global.NewPopover()
	r.Title = titleComponent()
	r.Selector = selectorPicker(r)

	group := r.Selector.SelectedElement()

	r.Average = NewAverage(group)
	r.Median = NewMedian(group)

	r.NotationTable = notationTable(r)

	// time.AfterFunc(2*time.Second, func() {
	// 	r.popover.Show(r.PassageOrderTable)
	// })
}

func (n *NotationPage) Render() app.UI {
	return app.Div().Body(
		&nav_bar.NavBar{},

		main_content.NewMain().Body(
			app.Script().Src("https://unpkg.com/ionicons@7.1.0/dist/ionicons/ionicons.js"),

			n.Popover,
			n.Title,
			notationSettingsHeader(),
			notationSettings(n),
			notationGradeHeader(),
			GradeSummary(n),
			n.NotationTable,
		),
	)
}
