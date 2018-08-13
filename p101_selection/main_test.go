package p101

import (
	"testing"
	"fmt"
	"sort"
	"math/rand"
)

func generateArray(l int) []int {
	a := make([]int, l)

	for i, _ := range a {
		a[i] = rand.Intn(1000000)
	}

	return a
}
func TestWorks(t *testing.T) {
	rows := []struct {
		array []int
		index int
	}{
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{9, 9, 9, 9, 9, 9, 9}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6},
		{generateArray(4), 2},
		{generateArray(100), 15},
		{generateArray(1000), 123},
		{generateArray(100001), 95123},
	}

	for i, row := range rows {
		ac := compute(row.array, row.index)
		//ac2 := compute(row.array, row.index)
		ac2 := computeC(row.array, row.index, 7)
		sort.Ints(row.array)
		ex := row.array[row.index]

		if ac != ex || ac2 != ex {
			fmt.Printf("%+v: ac: %+v, ac2: %+v, ex: %+v\n", i, ac, ac2, ex)
			t.Fail()
		}
	}
}

var num = 100000000
var array = generateArray(num)

func BenchmarkCompute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compute(array, 95123)
	}
}

func BenchmarkComputeC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		computeC(array, 95123, 10)
	}
}
