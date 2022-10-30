package main

import "fmt"

func shuffleSort(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		left++
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--
	}
	return arr
}

func main() {
	arr := []int{11, 14, 3, 8, 18, 17, 43}
	res := shuffleSort(arr)
	fmt.Println(res)
}
