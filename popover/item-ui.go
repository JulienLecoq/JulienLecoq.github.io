package popover

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	mapp "juliano.com/passage-order/map"
	"juliano.com/passage-order/theme"
)

var itemStyle = map[string]string{
	"padding": theme.ItemPadding.String(),
	"cursor":  "pointer",
}

var notFirstItemStyle = mapp.Merge(itemStyle, map[string]string{
	"border-top": "1px solid Gainsboro",
	// "border-top": "1px solid WhiteSmoke",
})

func (p *Popover) itemsUi() app.UI {
	return app.Range(p.Items).Slice(func(index int) app.UI {
		item := p.Items[index]
		style := p.itemStyle(index)

		return app.Div().
			Text(item).
			Styles(mapp.Merge(style, item.Styles())).
			Class("popover-item").
			OnClick(func(ctx app.Context, e app.Event) {
				item.OnClick(ctx, e)
				p.Hide()
			})
	})
}

func (p *Popover) itemStyle(index int) map[string]string {
	if index == 0 {
		return itemStyle
	}

	return notFirstItemStyle
}
