package p27

import (
	"math/big"
	"github.com/mmmpa/parallel/common"
)

func compute(rectCount int) float64 {
	sum := 0.0
	width := 1.0 / float64(rectCount)

	for i := 0; i < rectCount; i++ {
		mid := (float64(i) + 0.5) * width
		height := 4.0 / (1.0 + mid*mid)
		sum += height
	}

	return sum * width
}

func bf(n float64) *big.Float {
	return big.NewFloat(n)
}

func bfi(n int) *big.Float {
	return big.NewFloat(float64(n))
}

func add(a, b *big.Float) *big.Float {
	return new(big.Float).Add(a, b)
}

func div(a, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}

func mul(a, b *big.Float) *big.Float {
	return new(big.Float).Mul(a, b)
}

func computeBig(rectCount int) *big.Float {
	sum := bf(0.0)
	width := div(bf(1.0), bfi(rectCount))

	bf05 := bf(0.5)
	bf40 := bf(4.0)
	bf10 := bf(1.0)

	for i := 0; i < rectCount; i++ {
		a := add(bfi(i), bf05)
		mid := mul(a, width)
		height := div(
			bf40,
			add(
				bf10,
				mul(mid, mid),
			),
		)
		sum = add(sum, height)
	}

	return mul(sum, width)
}

func computeC(rectCount, workers int) float64 {
	sum := 0.0
	width := 1.0 / float64(rectCount)

	ch := make(chan float64)
	works := common.SplitWorks(rectCount, workers)
	head := 0

	for _, n := range works {
		tail := head + n

		go func(ch chan float64, head, tail int) {
			sum := 0.0

			for i := head; i < tail; i++ {
				mid := (float64(i) + 0.5) * width
				height := 4.0 / (1.0 + mid*mid)
				sum += height
			}

			ch <- sum
		}(ch, head, tail)

		head = tail
	}

	for i := 0; i < workers; i++ {
		sum += <-ch
	}

	return sum * width
}

type WorkerResult struct {
	Head    int
	Heights []float64
}

func computeCC(rectCount, workers int) float64 {
	sum := 0.0
	width := 1.0 / float64(rectCount)

	ch := make(chan WorkerResult)
	works := splitWorks(rectCount, workers)
	head := 0

	for _, n := range works {
		tail := head + n

		go func(ch chan WorkerResult, head, tail int) {
			result := WorkerResult{
				Head:    head,
				Heights: make([]float64, tail-head),
			}

			for i := head; i < tail; i++ {
				mid := (float64(i) + 0.5) * width
				height := 4.0 / (1.0 + mid*mid)

				result.Heights[i-head] = height
			}

			ch <- result
		}(ch, head, tail)

		head = tail
	}

	heights := make([]float64, rectCount)

	for i := 0; i < workers; i++ {
		ms := <-ch
		for i, height := range ms.Heights {
			heights[ms.Head+i] = height
		}
	}

	for _, height := range heights {
		sum += height
	}

	return sum * width
}
