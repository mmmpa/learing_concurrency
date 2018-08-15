package p27

func compute(array []int) int {
	sum := 0

	for _, n := range array {
		sum += n
	}

	return sum
}

func computeC(array []int, workers int) int {
	partsSum := firstC(array, workers)
	sums := secondC(partsSum, workers)

	return sums[0]
}

func firstC(array []int, workers int) []int {
	ch := make(chan interface{})
	length := len(array)

	if workers > length {
		return array
	}
	works := workers
	partsSum := make([]int, works)

	for i := 0; i < workers; i++ {
		go func(ch chan interface{}, offset int) {
			sum := 0
			for i := offset; i < length; i += workers {
				sum += array[i]
			}

			partsSum[offset] = sum

			ch <- struct{}{}
		}(ch, i)
	}

	for i := 0; i < works; i++ {
		<-ch
	}

	return partsSum
}

func secondC(array []int, workers int) []int {
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
				array[offset] += array[offset+next]

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

	return array
}
