package notation

import (
	"juliano.com/passage-order/notation/graded_user"
	"juliano.com/passage-order/selector"
)

func selectorPicker(r *NotationPage) *selector.Selector[*graded_user.GradedGroup] {
	return &selector.Selector[*graded_user.GradedGroup]{
		SelectorId: "group-select",
		SelectName: "group",
		Label:      "Sélectionne un groupe : ",
		Elements:   graded_user.Groups,
		OnChange: func(g *graded_user.GradedGroup) {
			r.NotationTable.Group = g
			r.NotationTable.Update()

			r.NotationTable.Group = g
			r.Average.Update()
		},
	}
}
