package p131

import (
	"testing"
	"math/rand"
	"fmt"
)

func gen(l int) []int {
	return shuffle(generateArray(l))
}

func generateArray(l int) []int {
	a := make([]int, l)

	for i, _ := range a {
		a[i] = i
	}

	return a
}

func shuffle(data []int) []int {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}

	return data
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

	for i := 0; i < 1; i++ {
		rows := []struct {
			array  []int
			target int
			result int
		}{
			{generateArray(16), 1, 1},
			{generateArray(16), 9, 9},
			{generateArray(16), 16, -1},
			{generateArray(15), 14, 14},
			{generateArray(1000000), 100000, 100000},
		}

		for _, row := range rows {
			ex := compute(row.array, row.target, 4)
			ac := computeC(row.array, row.target, 4)

			if ex != row.result || ac != row.result {
				fmt.Printf("%d: %+v %+v \n", i, ex, ac)
				t.Fail()
			}
		}
	}
}

var num = 10000000
var arrays [][]int

func BenchmarkCompute(b *testing.B) {
	b.ResetTimer()

	arrays = make([][]int, b.N)
	for i := 0; i < b.N; i++ {
		arrays[i] = gen(num)
	}

	for i := 0; i < b.N; i++ {
		compute(arrays[i], rand.Intn(num), 4)
	}
}

func BenchmarkComputeC(b *testing.B) {
	b.ResetTimer()

	arrays = make([][]int, b.N)
	for i := 0; i < b.N; i++ {
		arrays[i] = gen(num)
	}
	for i := 0; i < b.N; i++ {
		computeC(arrays[i], rand.Intn(num), 8)
	}
}

