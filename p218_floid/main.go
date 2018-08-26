package p131

import (
	"math"
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

func compute(distances [][]float64) [][]float64 {
	nodes := len(distances)
	results := copyDistances(distances)

	for k := 0; k < nodes; k++ {
		for i := 0; i < nodes; i++ {
			for j := 0; j < nodes; j++ {
				results[i][j] = math.Min(results[i][j], results[i][k]+results[k][j])
			}
		}
	}

	return results
}

func computeC(distances [][]float64, workers int) [][]float64 {
	nodes := len(distances)
	results := copyDistances(distances)

	ch := make(chan interface{})

	for k := 0; k < nodes; k++ {
		for w := 0; w < workers; w++ {
			go func(w int) {
				for i := w; i < nodes; i += workers {
					for j := 0; j < nodes; j++ {
						results[i][j] = math.Min(results[i][j], results[i][k]+results[k][j])
					}
				}
				ch <- struct{}{}
			}(w)
		}

		for ww := 0; ww < workers; ww++ {
			<-ch
		}
	}

	return results
}
