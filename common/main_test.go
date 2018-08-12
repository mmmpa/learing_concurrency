package common

import (
	"testing"
	"fmt"
)

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
		total   int
		workers int
		works   []int
	}{
		{10, 1, []int{10}},
		{10, 2, []int{5, 5}},
		{10, 3, []int{4, 3, 3}},
		{10, 4, []int{3, 3, 2, 2}},
	}

	for _, row := range rows {
		r := splitWorks(row.total, row.workers)
		if !eq(r, row.works) {
			fmt.Printf("%+v %+v\n", r, row.works)
			t.Fail()
		}
	}
}
