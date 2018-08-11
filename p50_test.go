package main

import (
	"testing"
	"fmt"
)

var rectCount = 10000000
var workers = 4

func TestCompute(t *testing.T) {
	fmt.Println(compute(rectCount))
	fmt.Println(computeC(rectCount, workers))
	fmt.Println(computeCC(rectCount, workers))
}

func BenchmarkCompute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compute(rectCount)
	}
}

func BenchmarkComputeC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		computeC(rectCount, workers)
	}
}

func BenchmarkComputeCC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		computeCC(rectCount, workers)
	}
}
