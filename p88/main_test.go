package p88

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

func TestWorks(t *testing.T) {
	rows := []struct {
		Array []int
	}{
		{generateArray(10)},
		{generateArray(100)},
		{generateArray(1000)},
	}

	for _, row := range rows {
		ac := compute(row.Array)
		ex := computeC(row.Array, 7)
		ex2 := computeCC(row.Array, 7)

		if ex != ac || ex2 != ac {
			fmt.Printf("%+v %+v %+v\n", ac, ex, ex2)
			t.Fail()
		}
	}
}
