package root

import (
	"juliano.com/passage-order/popover"
)

func popoverComponent(r *Root) *popover.Popover {
	return &popover.Popover{
		Visible:             false,
		HideOnDocumentClick: true,
	}
}
