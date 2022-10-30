package main

import "fmt"

// Factor for calculate step
const factor = 1.247

func combSort(arr []int) []int {
	end := len(arr) - 1
	for end >= 1 {
		for i := 0; i+end < len(arr); i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		end = int(float64(end) / factor)
	}
	return arr
}

func main() {
	arr := []int{11, 14, 3, 8, 18, 17, 43}
	res := combSort(arr)
	fmt.Println(res)
}
