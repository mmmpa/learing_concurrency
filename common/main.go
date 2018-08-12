package common

func SplitWorks(total int, workers int) []int {
	rest := total%workers - 1
	base := total / workers
	works := make([]int, workers)

	for i := 0; i < workers; i++ {
		if rest < i {
			works[i] = base
		} else {
			works[i] = base + 1
		}
	}

	return works
}
