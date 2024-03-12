package time_picker

import (
	"fmt"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type TimePicker struct {
	app.Compo

	InputId   string
	InputName string
	Label     string

	value    string
	Time     time.Time
	OnChange func(time time.Time)
}

func (t *TimePicker) Render() app.UI {
	return app.
		Div().
		Body(
			app.Label().For(t.InputId).Text(t.Label),
			app.Input().ID(t.InputId).Name(t.InputName).Type("time").Value(t.Time.Format("15:04")).OnInput(t.onTimeChange).Required(true),
		)
}

func (t *TimePicker) onTimeChange(ctx app.Context, e app.Event) {
	t.value = ctx.JSSrc().Get("value").String()

	if t.value == "" {
		return
	}

	layout := "15:04" // Define the layout to match the format of the input string

	parsedTime, err := time.Parse(layout, t.value)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	t.Time = parsedTime

	if t.OnChange != nil {
		t.OnChange(parsedTime)
	}
}

func DefaultTime() time.Time {
	return time.Date(0, 0, 0, 8, 0, 0, 0, time.Local)
}
