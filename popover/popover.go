package popover

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/component"
	"juliano.com/passage-order/mapp"
	"juliano.com/passage-order/theme"
)

var style = map[string]string{
	"border-radius":    theme.SmallRadius.String(),
	"overflow":         "hidden",
	"box-shadow":       "rgba(99, 99, 99, 0.2) 0px 2px 8px 0px",
	"width":            "fit-content",
	"position":         "absolute",
	"z-index":          "100000000",
	"background-color": "white",
}

type Item interface {
	fmt.Stringer
	OnClick(ctx app.Context, e app.Event)
	Styles() map[string]string
}

type Popover struct {
	app.Compo

	Visible             bool
	HideOnDocumentClick bool
	Items               []Item
	Reference           app.UI
	styles              map[string]string
}

func (p *Popover) OnInit() {
	app.Window().AddEventListener("click", func(ctx app.Context, e app.Event) {
		if p.Visible && p.HideOnDocumentClick {
			p.Hide()
		}
	})
}

func (p *Popover) Render() app.UI {
	items := p.itemsUi()

	if !p.Visible {
		return app.Div().Style("display", "none")
	}

	// return app.Div().Body(items).Styles(style).Styles(p.styles)
	return app.Div().Body(
		// arrow(),
		items,
	).Styles(mapp.Merge(style, p.styles))
}

func (p *Popover) Hide() {
	p.Visible = false
	p.Update()
}

// reference should not be set to nil.
func (p *Popover) Show(reference app.UI, items []Item) {
	if reference != nil {
		p.Reference = reference
	}

	if items != nil {
		p.Items = items
	}

	bounding := component.BoundingClientRect(p.Reference)

	p.styles = map[string]string{
		"top":  bounding.Bottom.String(),
		"left": bounding.Left.String(),
	}

	p.Visible = true
	p.Update()
}

func NewPopover() *Popover {
	return &Popover{
		Visible:             false,
		HideOnDocumentClick: true,
	}
}
