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
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6},
		{generateArray(4), 2},
		{generateArray(100), 15},
		{generateArray(1000), 123},
		{generateArray(100000), 95123},
	}

	for i, row := range rows {
		ac := compute(row.array, row.index)
		sort.Ints(row.array)
		ex := row.array[row.index]

		if ac != ex {
			fmt.Printf("%+v: re: %+v, ex: %+v\n", i, ac, ex)
			t.Fail()
		}
	}
}
