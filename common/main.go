package common

func SplitWorks(total int, workers int) []int {
	rest := total%workers - 1
	base := total / workers

	w := workers
	if total <= workers {
		w = total
	}

	works := make([]int, w)

	for i := 0; i < w; i++ {
		if total <= workers {
			works[i] = 1
		} else if rest < i {
			works[i] = base
		} else {
			works[i] = base + 1
		}
	}

	return works
}
