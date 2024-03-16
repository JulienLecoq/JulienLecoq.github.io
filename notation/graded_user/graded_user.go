package graded_user

import (
	"fmt"

	"juliano.com/passage-order/user"
)

type GradedUser struct {
	user.User
	Grade float64
}

func (g *GradedUser) GradeFormatted() string {
	fmt.Printf("Grade: %.1f\n", g.Grade)

	// Check if grade is an integer
	if g.Grade == float64(int(g.Grade)) {
		// If so, format without decimal part
		return fmt.Sprintf("%.0f", g.Grade)
	}
	// Otherwise, format with one decimal place
	return fmt.Sprintf("%.1f", g.Grade)
}
