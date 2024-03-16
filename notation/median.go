package notation

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/notation/graded_user"
)

type Median struct {
	app.Compo

	Group *graded_user.GradedGroup
}

func (a *Median) Render() app.UI {
	text := fmt.Sprintf("Mediane : %.1f", a.Group.Median())
	return app.Div().Text(text)
}

func NewMedian(group *graded_user.GradedGroup) *Median {
	return &Median{
		Group: group,
	}
}
