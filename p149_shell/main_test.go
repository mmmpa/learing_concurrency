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
			fmt.Println(i, an, b[i])
			return false
		}
	}
	return true
}
func TestCompute(t *testing.T) {
	rand.Seed(1)

	for i := 0; i < 100; i++ {
		rows := [][]int{
			{9, 15, 18, 18, 0, 3},
			generateArray(rand.Intn(5000)),
			generateArray(rand.Intn(5000)),
			generateArray(rand.Intn(5000)),
			generateArray(rand.Intn(5000)),
			generateArray(rand.Intn(5000)),
			generateArray(rand.Intn(5000)),
			generateArray(rand.Intn(5000)),
		}

		for i, row := range rows {
			ex := clone(row)
			sort.Ints(ex)

			ac1 := compute(clone(row))
			ac2 := computeCC(clone(row), 7)

			if !eq(ex, ac1) || !eq(ex, ac2) {
				fmt.Printf("%d: %+v %+v\n", i, eq(ex, ac1), eq(ex, ac2))
				fmt.Printf("%d: %+v %+v\n", i, eq(ex, ac1), eq(ex, compute(ac2)))
				t.Fail()
			}
		}
	}
}

var num = 100000
var array = generateArray(num)

func BenchmarkInsert(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		insert(clone(array))
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
