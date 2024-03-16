package passage_order

import (
	"juliano.com/passage-order/time_picker"
	"time"
)

func timePicker(r *PassageOrderPage) *time_picker.TimePicker {
	return &time_picker.TimePicker{
		InputId:   "date-picker",
		InputName: "date-picker",
		Label:     "Sélectionne le début de l'heure de passage : ",
		Time:      time_picker.DefaultTime(),
		OnChange: func(time time.Time) {
			r.PassageOrderTable.StartTime = time
			r.PassageOrderTable.Update()
		},
	}
}
