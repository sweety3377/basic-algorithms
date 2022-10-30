package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) <= 2 {
		return arr
	}

	left := 0
	right := len(arr) - 1

	pivotIndex := rand.Int() % len(arr)

	arr[pivotIndex], arr[right] = arr[pivotIndex], arr[right]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{11, 14, 3, 8, 18, 17, 43}
	quickSort(arr)
	fmt.Println(arr)
}
