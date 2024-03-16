package mathh

import "sort"

func Median(values []float64) float64 {
	sort.Float64s(values) // Sort the slice.
	n := len(values)

	// If the length of the slice is odd, return the middle element.
	if n%2 != 0 {
		return values[n/2]
	}

	// If the length of the slice is even, return the average of the two middle elements.
	return (values[(n/2)-1] + values[n/2]) / 2.0

}
