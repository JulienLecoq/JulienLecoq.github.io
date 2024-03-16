package passage_order

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/global"
	"juliano.com/passage-order/main_content"
	"juliano.com/passage-order/nav_bar"
	"juliano.com/passage-order/popover"

	"juliano.com/passage-order/passage_duration"
	"juliano.com/passage-order/passage_order_table"
	"juliano.com/passage-order/selector"
	"juliano.com/passage-order/time_picker"
	"juliano.com/passage-order/title"
	"juliano.com/passage-order/user"
)

type PassageOrderPage struct {
	app.Compo

	Popover           *popover.Popover
	Title             *title.Title
	Selector          *selector.Selector[user.Group]
	TimePicker        *time_picker.TimePicker
	PassageDuration   *passage_duration.PassageDuration
	PassageOrderTable *passage_order_table.PassageOrderTable
}

func (r *PassageOrderPage) OnInit() {
	r.Title = titleComponent()
	r.Selector = selectorPicker(r)
	r.TimePicker = timePicker(r)
	r.PassageDuration = passageDurationPicker(r)
	r.PassageOrderTable = passageOrderTable(r)
	// r.Popover = global.Popover
	r.Popover = global.NewPopover()

	// time.AfterFunc(2*time.Second, func() {
	// 	r.popover.Show(r.PassageOrderTable)
	// })
}

func (r *PassageOrderPage) Render() app.UI {
	return app.Div().Body(
		&nav_bar.NavBar{},

		main_content.NewMain().Body(
			app.Script().Src("https://unpkg.com/ionicons@7.1.0/dist/ionicons/ionicons.js"),

			r.Popover,
			r.Title,
			passageSettingsHeader(),
			passageSettings(r),
			passageOrderHeader(),
			r.PassageOrderTable,
		),
	)
}

func (r *PassageOrderPage) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
