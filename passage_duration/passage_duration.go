package passage_duration

import (
	"strconv"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PassageDuration struct {
	app.Compo

	InputId   string
	InputName string
	Label     string

	Duration time.Duration
	OnChange func(duration time.Duration)
}

func (p *PassageDuration) Render() app.UI {
	return app.
		Div().
		Body(
			app.Label().For(p.InputId).Text(p.Label),
			app.Input().ID(p.InputId).Name(p.InputName).Type("number").OnInput(p.onDurationChange).Value(int(p.Duration.Minutes())).Min(1).Required(true),
		)
}

func (p *PassageDuration) emitDefault() {
	p.Duration = time.Duration(10) * time.Minute
	p.OnChange(p.Duration)
	p.Update()
}

func (p *PassageDuration) onDurationChange(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value").String()
	duration, err := strconv.ParseUint(value, 10, 64)

	if duration == 0 || err != nil {
		p.emitDefault()
		return
	}

	p.Duration = time.Duration(duration) * time.Minute

	if p.OnChange != nil {
		p.OnChange(p.Duration)
	}
}
