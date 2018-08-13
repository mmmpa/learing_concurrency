package p88

import "github.com/mmmpa/parallel/common"

func compute(array []int) int {
	sum := 0

	for _, n := range array {
		sum += n
	}

	return sum
}

type SumResult struct {
	Index int
	Sum   int
}

func computeC(array []int, workers int) int {
	sum := 0

	ch := make(chan SumResult)
	works := common.SplitWorks(len(array), workers)
	head := 0

	for i, n := range works {
		tail := head + n

		go func(ch chan SumResult, i, head, tail int) {
			sum := 0

			for i := head; i < tail; i++ {
				sum += array[i]
			}

			ch <- SumResult{
				Index: i,
				Sum:   sum,
			}
		}(ch, i, head, tail)

		head = tail
	}

	results := make([]int, len(works))

	for i := 0; i < workers; i++ {
		result := <-ch
		results[result.Index] = result.Sum
	}

	for _, n := range results {
		sum += n
	}

	return sum
}

func computeCC(array []int, workers int) int {
	sum := 0

	ch := make(chan SumResult)
	works := common.SplitWorks(len(array), workers)

	for i, _ := range works {
		go func(ch chan SumResult, offset int) {
			sum := 0

			for i := offset; i < len(array); i+=workers {
				sum += array[i]
			}

			ch <- SumResult{
				Index: offset,
				Sum:   sum,
			}
		}(ch, i)
	}

	results := make([]int, len(works))

	for i := 0; i < workers; i++ {
		result := <-ch
		results[result.Index] = result.Sum
	}

	for _, n := range results {
		sum += n
	}

	return sum
}
