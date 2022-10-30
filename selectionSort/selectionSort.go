package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := 1; i != len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort(arr []int) []int {
	newArr := make([]int, 0, len(arr))
	arrLen := len(arr)

	byte
	var smallest int
	for i := 0; i != arrLen; i++ {
		smallest = findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}

	return newArr
}

func main() {
	arr := []int{5, 3, 6, 2, 10}
	newArr := selectionSort(arr)
	fmt.Println(newArr)
}
