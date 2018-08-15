package p27

import (
	"testing"
	"fmt"
	"math/rand"
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

func TestCompute(t *testing.T) {
	rows := [][]int{
		generateArray(6),
		generateArray(11),
		generateArray(100),
		generateArray(7777),
	}

	for i, row := range rows {
		ex := compute(clone(row))
		ac1 := computeC(clone(row), 1)
		ac2 := computeC(clone(row), 7)

		if ex != ac1 || ex != ac2 {
			fmt.Printf("%d: %+v %+v %+v\n", i, ex, ac1, ac2)
			t.Fail()
		}
	}
}
