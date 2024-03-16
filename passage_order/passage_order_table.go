package passage_order

import (
	"time"

	"juliano.com/passage-order/passage_order_table"
)

func passageOrderTable(r *PassageOrderPage) *passage_order_table.PassageOrderTable {
	return &passage_order_table.PassageOrderTable{
		Group:           r.Selector.SelectedElement(),
		StartTime:       r.TimePicker.Time,
		PassageDuration: 12 * time.Minute,
	}
}
