package p131

import (
	"testing"
	"fmt"
	"math/rand"
	"sort"
)

func generateArray(l int) []int {
	a := make([]int, l)

	for i, _ := range a {
		a[i] = rand.Intn(100000)
	}

	return a
}

func clone(base []int) []int {
	a := make([]int, len(base))

	for i, _ := range a {
		a[i] = base[i]
	}

	return a
}

func eq(a []int, b []int) bool {
	for i, an := range a {
		if an != b[i] {
			return false
		}
	}
	return true
}
func TestCompute(t *testing.T) {
	rows := [][]int{
		generateArray(6),
		generateArray(11),
		generateArray(16),
		generateArray(24),
		generateArray(100),
		generateArray(1000),
		generateArray(3000),
		generateArray(100),
		generateArray(1000),
		generateArray(100),
		generateArray(500),
	}

	for i, row := range rows {
		ex := clone(row)
		sort.Ints(ex)

		ac1 := compute(clone(row))
		ac2 := computeC(clone(row), 3)

		if !eq(ex, ac1) || !eq(ex, ac2) {
			fmt.Printf("%d: %+v %+v %+v\n", i, ex, ac1, ac2)
			t.Fail()
		}
	}
}

var num = 30000
var array = generateArray(num)

func BenchmarkCompute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compute(clone(array))
	}
}

func BenchmarkComputeC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		computeC(clone(array), 1)
	}
}
