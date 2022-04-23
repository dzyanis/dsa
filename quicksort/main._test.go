package main

import "testing"

func BenchmarkQuicksort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Quicksort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	}
}
