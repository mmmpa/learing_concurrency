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

func printer2(ds [][]int) {
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
				{max,   7,   2,   5, max},
				{  7, max, max,   3,   8},
				{  2, max, max,   4,   3},
				{  5,   3,   4, max,   2},
				{max,   8,   3,   2, max},
			},
			//generate(1000),
		}

		fmt.Println("prepared")
		runtime.GC()

		for _, row := range rows {
			ac := compute(row)
			ac3 := computeC(row, 4)
			println("start ■")
			//itemPrinter(ac)
			printer(row)
			printer2(ac)
			fmt.Println("")
			printer2(ac3)
			//itemPrinter(ac3)
			println("end ■")
			// fmt.Println(ac2)
			//printer(row)
			// ac := computeC(row.array, row.target, 4)

		}
	}
}
