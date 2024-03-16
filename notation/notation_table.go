package notation

import (
	// "fmt"

	"juliano.com/passage-order/notation/graded_user"
	"juliano.com/passage-order/notation/notation_table"
)

func notationTable(r *NotationPage) *notation_table.NotationTable {
	group := r.Selector.SelectedElement()

	return &notation_table.NotationTable{
		Group: group,
		OnChange: func(user *graded_user.GradedUser, prevGrade float64, newGrade float64) {
			// fmt.Printf("User grade updated!\n")

			r.Average.Update()
			r.Median.Update()
		},
	}
}
