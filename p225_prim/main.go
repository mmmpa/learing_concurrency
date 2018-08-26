package p131

import (
	"math"
	"fmt"
)

func copyDistances(distances [][]float64) [][]float64 {
	nodes := len(distances)

	ds := make([][]float64, nodes)

	for i, _ := range ds {
		ds[i] = make([]float64, nodes)
		copy(ds[i], distances[i])
	}

	return ds
}

func compute(distances [][]float64) [][]int {
	nodes := len(distances)

	nearNode := make([]int, nodes)
	minDist := make([]float64, nodes)
	result := make([][]int, nodes)

	for i := 0; i < nodes; i++ {
		nearNode[i] = 0
		minDist[i] = distances[i][0]
		result[i] = make([]int, 2)
	}

	k := -1
	for i := 0; i < nodes-1; i++ {
		min := float64(math.MaxInt16)

		fmt.Println("i", i)

		//
		// 現在隣接とされているノードの中で最短距離にあるノードを調べる。
		// j == 0 は初期ノードなので除外する
		// 現在隣接とされているノードは
		// - 初期状態では最初のノード
		// - 次回からは前段のループの最終項で更新される
		// - 既に訪れたノードは -1 になるため除外される
		for j := 1; j < nodes; j++ {
			if 0 <= minDist[j] && minDist[j] <= min {
				min = minDist[j]
				k = j
			}
		}
		fmt.Println(minDist)
		fmt.Println("k", k)

		result[i][0] = nearNode[k]
		result[i][1] = k
		minDist[k] = -1

		// 最短距離にあるノードからの距離をはかり、
		// より近ければ上書きして隣接ノードをその最短距離ノードとする
		for j := 0; j < nodes; j++ {
			if distances[j][k] < minDist[j] {
				minDist[j] = distances[j][k]
				nearNode[j] = k
			}
		}
		fmt.Println(minDist)
		fmt.Println(nearNode)
		fmt.Println("")

	}

	return result
}

func computeC(distances [][]float64, workers int) [][]int {
	nodes := len(distances)
	ch := make(chan interface{})

	nearNode := make([]int, nodes)
	minDist := make([]float64, nodes)
	result := make([][]int, nodes)

	for i := 0; i < nodes; i++ {
		nearNode[i] = 0
		minDist[i] = distances[i][0]
		result[i] = make([]int, 2)
	}

	for i := 0; i < nodes-1; i++ {
		mins := make([]float64, workers)
		for w := 0; w < workers; w++ {
			mins[w] = float64(math.MaxInt16)
		}
		ks := make([]int, workers)

		fmt.Println("i", i)

		for w := 0; w < workers; w++ {
			go func(w int) {
				for j := w + 1; j < nodes; j += workers {
					if 0 <= minDist[j] && minDist[j] <= mins[w] {
						mins[w] = minDist[j]
						ks[w] = j
					}
				}
				ch <- struct{}{}
			}(w)
		}

		for ww := 0; ww < workers; ww++ {
			<-ch
		}

		k := secondC(mins, ks, workers)
		fmt.Println(mins, ks, k)

		result[i][0] = nearNode[k]
		result[i][1] = k
		minDist[k] = -1

		// 最短距離にあるノードからの距離をはかり、
		// より近ければ上書きして隣接ノードをその最短距離ノードとする
		for w := 0; w < workers; w++ {
			go func(w int) {
				for j := w; j < nodes; j += workers {
					if distances[j][k] < minDist[j] {
						minDist[j] = distances[j][k]
						nearNode[j] = k
					}
				}
				ch <- struct{}{}
			}(w)
		}

		for ww := 0; ww < workers; ww++ {
			<-ch
		}
	}

	return result
}
func secondC(array []float64, ks []int, workers int) int {
	ch := make(chan interface{})
	length := len(array)
	step := 2
	next := 1

	for {
		nextWorkers := 0

		for offset := 0; offset < workers; offset += step {
			if offset+next >= length {
				continue
			}

			nextWorkers++
			go func(ch chan interface{}, offset, next int) {
				if array[offset] < array[offset+next] {
					array[offset] = array[offset]
					ks[offset] = ks[offset]
				} else {
					array[offset] = array[offset+next]
					ks[offset] = ks[offset+next]
				}

				ch <- struct{}{}
			}(ch, offset, next)
		}

		if nextWorkers == 0 {
			break
		}

		for i := 0; i < nextWorkers; i++ {
			<-ch
		}

		step *= 2
		next *= 2
	}

	return ks[0]
}
