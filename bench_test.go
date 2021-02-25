package main

import "testing"

func BenchmarkSliceAppend(b *testing.B) {
	is := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is = append(is, i)
	}
}

func BenchmarkSliceAppendAlocOnce(b *testing.B) {
	is := make([]int, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		is[i] = i
	}
}
