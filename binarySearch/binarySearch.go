package main

import (
	"fmt"
)

func binarySearch(arr []int, item int) int {
	// Start index
	start := 0

	// Last index
	end := len(arr) - 1

	// For end index higher than start index
	for start <= end {
		// Mid number index
		midIndex := (start + end) / 2

		// Mid number
		midNum := arr[midIndex]

		// Check
		if midNum == item {
			return midIndex
		}

		// If number higher than item
		// End index decrement
		if midNum > item {
			end = midIndex - 1
		} else {
			// Else mid number lowest than item
			// Start index decrement
			start = midIndex + 1
		}
	}

	// Return 0 if not exist
	return 0
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 10, 12, 13, 15, 17, 20}
	res := binarySearch(arr, 17)
	fmt.Println(res)
}