package p27

import (
	"testing"
	"fmt"
)

var rectCount = 10000000
var workers = 8

func eq(a []int, b []int) bool {
	for i, an := range a {
		if an != b[i] {
			return false
		}
	}
	return true
}

func TestCompute(t *testing.T) {
	rows := []struct {
		total   int
		workers int
		comp    func(int, int) float64
	}{
		{10000, 1, computeC},
		{10000, 1, computeCC},
		{10000, 2, computeCC},
		{10000, 10, computeCC},
	}

	for _, row := range rows {
		ex := compute(row.total)
		ac := row.comp(row.total, row.workers)
		if ex != ac {
			fmt.Printf("%+v, %+v: %+v %+v\n", row.total, row.workers, ex, ac)
			t.Fail()
		}
	}
}

func TestComputeC(t *testing.T) {
	ex := compute(100000)
	ac := computeCCC(100000, 1)

	if ex != ac {
		fmt.Printf("%+v %+v\n", ex, ac)
		t.Fail()
	}

	for i := 0; i < 100; i++ {
		exc := computeCCC(100000, 10)
		acc := computeCCC(100000, 10)

		if exc != acc {
			fmt.Printf("%+v %+v\n", exc, acc)
			t.Fail()
		}
	}
}

func BenchmarkCompute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compute(rectCount)
	}
}

func BenchmarkComputeBig(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		computeBig(rectCount)
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

func BenchmarkComputeCCC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		computeCCC(rectCount, workers)
	}
}
