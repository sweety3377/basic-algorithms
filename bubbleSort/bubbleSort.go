package main

import "fmt"

func bubbleSort(arr []int) []int {
	// For loop on all array
	for i := range arr {
		// For loop all array again
		// For check all pairs to replace
		for j := 0; j < len(arr)-i-1; j++ {
			// if arr[j] (for ex: 5) higher that arr[j+1] (for ex: 6)
			// We will replace 5 on 6 and sort again
			if arr[j] > arr[j+1] {
				// Swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{11, 14, 1, 2, 3, 8, 18, 17, 43}
	res := bubbleSort(arr)
	fmt.Println(res)
}
