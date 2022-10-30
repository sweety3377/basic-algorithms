package main

import "fmt"

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	return merge(mergeSort(slice[:mid]), mergeSort(slice[mid:]))
}

func merge(first, second []int) []int {
	size := len(first) + len(second)
	i, j, count := 0, 0, 0
	slice := make([]int, size, size)

	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			slice[count] = first[i]
			count++
			i++
		} else {
			slice[count] = second[j]
			count++
			j++
		}
	}

	for i < len(first) {
		slice[count] = first[i]
		count++
		i++
	}

	for j < len(second) {
		slice[count] = second[j]
		count++
		j++
	}

	return slice
}

func main() {
	arr := []int{11, 14, 3, 8, 18, 17, 43}
	res := mergeSort(arr)

	fmt.Println(res)
}
