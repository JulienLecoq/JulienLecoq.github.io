package root

import (
	"juliano.com/passage-order/passage_duration"
	"time"
)

func passageDurationPicker(r *Root) *passage_duration.PassageDuration {
	return &passage_duration.PassageDuration{
		InputId:   "passage-duration",
		InputName: "passage-duration",
		Label:     "Sélectionne la durée d'un passage : ",
		Duration:  time.Duration(12) * time.Minute,
		OnChange: func(duration time.Duration) {
			r.PassageOrderTable.PassageDuration = duration
			r.PassageOrderTable.Update()
		},
	}
}
