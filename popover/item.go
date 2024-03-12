package popover

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type PopoverItem struct {
	Name    string
	styles  map[string]string
	onClick app.EventHandler
}

func NewPopoverItem(name string, onClick app.EventHandler, styles map[string]string) *PopoverItem {
	return &PopoverItem{
		Name:    name,
		onClick: onClick,
		styles:  styles,
	}
}

func (p PopoverItem) String() string {
	return p.Name
}

func (p *PopoverItem) OnClick(ctx app.Context, event app.Event) {
	if p.onClick != nil {
		p.onClick(ctx, event)
	}
}

func (p PopoverItem) Styles() map[string]string {
	return p.styles
}
