package p131

import (
	"testing"
	"fmt"
	"math/rand"
	"sort"
	"time"
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
	rand.Seed(time.Now().Unix())

	rows := [][]int{
		{9, 15, 18, 18, 0, 3},
		generateArray(6),
		generateArray(11),
		generateArray(16),
		generateArray(24),
		generateArray(100),
		generateArray(500),
		generateArray(1000),
		generateArray(2000),
		generateArray(3100),
		generateArray(100),
		generateArray(1000),
		generateArray(100),
		generateArray(500),
	}

	for i, row := range rows {
		ex := clone(row)
		sort.Ints(ex)

		ac1 := compute(clone(row))
		ac2 := computeC(clone(row), 7)

		if !eq(ex, ac1) || !eq(ex, ac2) {
			fmt.Printf("%d: %+v %+v %+v %+v\n", i, row, ex, ac1, ac2)
			t.Fail()
		}
	}
}

var num = 10000
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
		computeC(clone(array), 6)
	}
}
