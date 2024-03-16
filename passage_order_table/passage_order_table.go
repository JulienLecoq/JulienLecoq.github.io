package passage_order_table

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/mapp"
	"juliano.com/passage-order/table"
	"juliano.com/passage-order/user"
)

type PassageOrderTable struct {
	app.Compo

	Group           user.Group
	StartTime       time.Time
	PassageDuration time.Duration

	Styles map[string]string
}

func (p *PassageOrderTable) Render() app.UI {
	TRs := append([]app.UI{p.headings()}, p.body()...)
	style := mapp.Merge(table.TableStyle, p.Styles)

	return app.Table().Body(TRs...).Styles(style)
}
