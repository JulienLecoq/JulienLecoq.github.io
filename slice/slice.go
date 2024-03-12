package slice

import (
	"math/rand"
)

func Shuffle[T any](slice []T) {
	// Shuffle the slice
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// Shuffle returns a new slice with the elements of the input slice shuffled.
func PureShuffle[T any](slice []T) []T {
	// Create a copy of the original slice to keep it unchanged
	shuffled := make([]T, len(slice))
	copy(shuffled, slice)

	// Shuffle the copied slice
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

func Delete[T any](slice []T, i int) []T {
	if i >= 0 && i < len(slice) {
		// Create a new slice to hold the result.
		result := make([]T, 0, len(slice)-1)
		// Append elements before index i.
		result = append(result, slice[:i]...)
		// Append elements after index i.
		result = append(result, slice[i+1:]...)
		return result
	}

	// Return the original slice if index is out of bounds.
	return slice

	// if i >= 0 && i < len(slice) {
	// 	// Append the elements before index i and elements after index i.
	// 	return append(slice[:i], slice[i+1:]...)
	// }
	//
	// // Return the original slice if index is out of bounds.
	// return slice
}
