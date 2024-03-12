package root

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/global"
	"juliano.com/passage-order/popover"

	// "juliano.com/passage-order/ionicon"
	"juliano.com/passage-order/passage_duration"
	"juliano.com/passage-order/passage_order_table"
	"juliano.com/passage-order/selector"
	"juliano.com/passage-order/theme"
	"juliano.com/passage-order/time_picker"
	"juliano.com/passage-order/title"
	"juliano.com/passage-order/user"
)

var rootStyle = map[string]string{
	"padding": theme.ContentMainMargin,
}

type Root struct {
	app.Compo

	Popover           *popover.Popover
	Title             *title.Title
	Selector          *selector.Selector[user.Group]
	TimePicker        *time_picker.TimePicker
	PassageDuration   *passage_duration.PassageDuration
	PassageOrderTable *passage_order_table.PassageOrderTable
}

func (r *Root) OnMount(ctx app.Context) {
	global.Popover = popoverComponent(r)

	r.Title = titleComponent()
	r.Selector = selectorPicker(r)
	r.TimePicker = timePicker(r)
	r.PassageDuration = passageDurationPicker(r)
	r.PassageOrderTable = passageOrderTable(r)
	r.Popover = global.Popover

	// time.AfterFunc(2*time.Second, func() {
	// 	r.popover.Show(r.PassageOrderTable)
	// })
}

func (r *Root) Render() app.UI {
	return app.Div().Body(
		app.Script().Type("module").Src("https://unpkg.com/ionicons@7.1.0/dist/ionicons/ionicons.esm.js"),

		r.Popover,
		r.Title,
		passageSettingsHeader(),
		passageSettings(r),
		passageOrderHeader(),
		r.PassageOrderTable,
	).Styles(rootStyle)
}

func (r *Root) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
