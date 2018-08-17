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
		a[i] = rand.Intn(1000000)
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
			fmt.Println(i, an, b[i])
			return false
		}
	}
	return true
}
func TestCompute(t *testing.T) {
	rand.Seed(1)

	for i := 0; i < 1000; i++ {
		rows := [][]int{
			generateArray(100),
			{9, 15, 18, 18, 0, 3},
			generateArray(rand.Intn(10)),
		}

		for i, row := range rows {
			ex := clone(row)
			sort.Ints(ex)

			ac1 := compute(clone(row))
			ac2 := computeC(clone(row), 2)

			if !eq(ex, ac1) || !eq(ex, ac2) {
				fmt.Printf("%d: %+v %+v %+v\n", i, ex, ac1, ac2)
				t.Fail()
			}
		}
	}
}

var num = 1000000
var array = generateArray(num)

func BenchmarkQ(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		computeQ(clone(array))
	}
}

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
