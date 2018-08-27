package p131

import (
	"testing"
	"math/rand"
	"fmt"
	"runtime/debug"
	"runtime"
	"math"
)

func eq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, an := range a {
		if an != b[i] {
			fmt.Printf("%v; %v %v %b %b", i, an, b[i], an, b[i])
			return false
		}
	}
	return true
}

func printer(ds [][]float64) {
	for _, d := range ds {
		fmt.Println(d)
	}
}

func TestCompute(t *testing.T) {
	debug.SetGCPercent(-1)
	rand.Seed(1)

	max := float64(math.MaxInt16)

	for i := 0; i < 1; i++ {
		rows := [][][]float64{
			{
				{0, 3, 1, max, max, max},
				{3, 0, 1, 4, 2, max},
				{1, 1, 0, max, 4, max},
				{max, 4, max, 0, 1, 3},
				{max, 2, 4, 1, 0, 2},
				{max, max, max, 3, 2, 0},
			},
			{
				{0,   max, max, max, max,   1},
				{max,   0, max, max,   1, max},
				{max, max,   0,   1, max,   1},
				{max, max,   1,   0,   1, max},
				{max,   1, max,   1,   0, max},
				{1,   max,   1, max, max,   0},
			},
			//generate(1000),
		}

		fmt.Println("prepared")
		runtime.GC()

		for _, row := range rows {
			ac := compute(row)
			ac3 := computeC(row, 8)
			println("start ■")
			//itemPrinter(ac)
			printer(row)
			printer(ac)
			fmt.Println("")
			printer(ac3)
			//itemPrinter(ac3)
			println("end ■")
			// fmt.Println(ac2)
			//printer(row)
			// ac := computeC(row.array, row.target, 4)

		}
	}
}
