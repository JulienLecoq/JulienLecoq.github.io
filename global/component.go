package global

import "juliano.com/passage-order/popover"

var Popover *popover.Popover

func NewPopover() *popover.Popover {
	Popover = popover.NewPopover()

	return Popover
}
