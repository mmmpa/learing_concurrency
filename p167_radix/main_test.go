package p131

import (
	"testing"
	"math/rand"
	"sort"
	"fmt"
)

func generateArray(l int) []int {
	a := make([]int, l)

	for i, _ := range a {
		a[i] = rand.Intn(20)
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
			fmt.Printf("%v; %v %v %b %b", i, an, b[i], an, b[i])
			return false
		}
	}
	return true
}
func TestCompute(t *testing.T) {
	rand.Seed(1)

	for i := 0; i < 100; i++ {
		rows := [][]int{
			{445, 425, 511},
			{5, 3, 1, 6},
			{445, 425, 511},
			generateArray(1000),
			generateArray(rand.Intn(10000) + 1),
		}

		for _, row := range rows {
			ex := clone(row)
			sort.Ints(ex)

			ac1 := compute(clone(row))
			//ac2 := compute(clone(row))
			ac2 := computeC(clone(row), 10)

			if !eq(ex, ac1) || !eq(ex, ac2) {
				fmt.Printf("%d: %+v %+v %+v\n", i, ex, ac1, ac2)
				t.Fail()
			}
		}
	}
}

var num = 1000000
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
		computeC(clone(array), 12)
	}
}

func BenchmarkComputeCC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		computeC(clone(array), 1)
	}
}
