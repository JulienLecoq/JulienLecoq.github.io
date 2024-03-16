package graded_user

import "juliano.com/passage-order/mathh"

func (g *GradedGroup) Average() float64 {
	sum := 0.0

	for _, user := range g.Users {
		// fmt.Printf("Grade from average: %d \n", user.Grade)
		sum += user.Grade
	}

	return sum / float64(len(g.Users))
}

func (g *GradedGroup) Median() float64 {
	grades := []float64{}

	for _, user := range g.Users {
		grades = append(grades, user.Grade)
	}

	return mathh.Median(grades)
}
