package p131

import (
	"testing"
	"math/rand"
	"fmt"
	"runtime/debug"
	"runtime"
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

func uniq(array []int) bool {
	set := make(map[int]bool)

	for _, n := range array {
		if n == -1 {
			continue
		}
		if set[n] {
			return false
		}

		set[n] = true
	}

	return true
}

func uniqI(array []*Item) bool {
	set := make(map[int]bool)

	for _, n := range array {
		if n == nil {
			continue
		}
		if set[n.N] {
			return false
		}

		set[n.N] = true
	}

	return true
}

func generate(nodes int) [][]int {
	adj := make([][]int, nodes)

	for i, _ := range adj {
		adj[i] = make([]int, nodes)
	}

	for i := 0; i < nodes; i++ {
		for j := i + 1; j < nodes; j++ {
			base := rand.Intn(10)
			tf := 0
			if base > 8 {
				tf = 1
			}

			adj[i][j] = tf
			for k := j; k < nodes; k++ {
				adj[j][i] = tf
			}
		}
	}

	return adj
}

func printer(adj [][]int) {
	for _, l := range adj {
		fmt.Println(l)
	}
}

func itemPrinter(items []*Item) {
	for _, l := range items {
		if l == nil || l.Path == nil {
			continue
		}

		if l.Path.Length > 0 {
			node := l.Path.Head

			fmt.Print("head ")
			for node != nil {
				fmt.Printf("%d ", node.N)
				node = node.Next
			}

			fmt.Print("\n")
		}
	}
}

func TestCompute(t *testing.T) {
	debug.SetGCPercent(-1)
	rand.Seed(1)

	for i := 0; i < 1; i++ {
		rows := [][][]int{
			// {
			// 	{0, 1, 1, 1, 0},
			// 	{1, 0, 1, 0, 1},
			// 	{1, 1, 0, 1, 1},
			// 	{1, 0, 1, 0, 1},
			// 	{0, 1, 1, 1, 0},
			// },
			generate(10),
			generate(15),
			generate(1000),
			//generate(1000),
		}

		fmt.Println("prepared")
		runtime.GC()

		for _, row := range rows {
			ac := compute(row)
			ac2 := computeA(row)
			ac3 := computeC(row, 8)
			println("start ■")
			//itemPrinter(ac)
			println("")
			//itemPrinter(ac3)
			println("end ■")
			// fmt.Println(ac2)
			//printer(row)
			// ac := computeC(row.array, row.target, 4)

			runtime.GC()
			if !uniqI(ac) || !uniq(ac2) || !uniqI(ac3) {
				fmt.Printf("%d: %+v \n", i, ac)
				t.Fail()
			}
		}
	}
}

var array = generate(3000)

func BenchmarkCompute(b *testing.B) {
	debug.SetGCPercent(-1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		compute(array)
	}
}

func BenchmarkComputeC(b *testing.B) {
	debug.SetGCPercent(-1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		computeC(array, 8)
	}
}
