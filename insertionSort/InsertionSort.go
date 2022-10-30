package main

import "fmt"

func insertionSort(arr []int) []int {
	// Start for loop
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		ind := i

		for ind > 0 && arr[ind-1] > num {
			arr[ind] = arr[ind-1]
			ind--
		}

		arr[ind] = num
	}

	return arr
}

func main() {
	arr := []int{11, 14, 3, 8, 18, 17, 43}
	res := insertionSort(arr)
	fmt.Println(res)
}
