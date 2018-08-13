package p93

import (
	"testing"
	"fmt"
	"math/rand"
)

func generateArray(l int) []int {
	a := make([]int, l)

	for i, _ := range a {
		a[i] = rand.Intn(100)
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

func TestWorks(t *testing.T) {
	rows := []struct {
		Array []int
	}{
		{generateArray(5)},
		{generateArray(100)},
		{generateArray(10000000)},
	}

	for _, row := range rows {
		ac := compute(row.Array)
		ex := computeC(row.Array, 10)

		if !eq(ex, ac) {
			fmt.Printf("%+v %+v\n", ac, ex)
			t.Fail()
		}
	}
}
