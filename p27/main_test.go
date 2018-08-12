package main

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

func TestWorks(t *testing.T) {
	rows := []struct {
		total   int
		workers int
		works   []int
	}{
		{10, 1, []int{10}},
		{10, 2, []int{5, 5}},
		{10, 3, []int{4, 3, 3}},
		{10, 4, []int{3, 3, 2, 2}},
	}

	for _, row := range rows {
		r := splitWorks(row.total, row.workers)
		if !eq(r, row.works) {
			fmt.Printf("%+v %+v\n", r, row.works)
			t.Fail()
		}
	}
}

func TestCompute(t *testing.T) {
	rows := []struct {
		total   int
		workers int
		comp func(int, int) float64
	}{
		{10000, 1, computeC},
		{10000, 1, computeCC},
		// {10000, 2, computeC},
		{10000, 2, computeCC},
		// {10000, 10, computeC},
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
