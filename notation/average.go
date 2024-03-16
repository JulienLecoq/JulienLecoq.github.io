package notation

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"juliano.com/passage-order/notation/graded_user"
)

type Average struct {
	app.Compo

	Group *graded_user.GradedGroup
}

func (a *Average) Render() app.UI {
	text := fmt.Sprintf("Moyenne : %.1f", a.Group.Average())
	return app.Div().Text(text)
}

func NewAverage(group *graded_user.GradedGroup) *Average {
	return &Average{Group: group}
}
