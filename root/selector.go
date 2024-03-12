package root

import (
	// "fmt"

	// "juliano.com/passage-order/selector"
	"juliano.com/passage-order/selector"
	"juliano.com/passage-order/user"
)

func selectorPicker(r *Root) *selector.Selector[user.Group] {
	return &selector.Selector[user.Group]{
		SelectorId: "group-select",
		SelectName: "group",
		Label:      "Sélectionne un groupe : ",
		Elements:   user.Groups,
		OnChange: func(g user.Group) {
			r.PassageOrderTable.Group = g
			r.PassageOrderTable.Update()
		},
	}
}
