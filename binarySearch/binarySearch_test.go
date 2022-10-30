package main

import (
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		numArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

		for _, num := range numArr {
			_ = binarySearch(numArr, num)
		}
	}
}
