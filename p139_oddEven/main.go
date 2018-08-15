package p131

func compute(array []int) []int {
	length := len(array)

	ex := true
	start := 0

	for ex || start == 1 {
		ex = false

		for i := start; i < length-1; i += 2 {
			if array[i] > array[i+1] {
				t := array[i]
				array[i] = array[i+1]
				array[i+1] = t
				ex = true
			}
		}

		if start == 0 {
			start = 1
		} else {
			start = 0
		}
	}

	return array
}

func computeC(array []int, workers int) []int {
	length := len(array)
	ch := make(chan bool)

	works := length/workers + 1
	if works%2 != 0 {
		works++
	}

	ex := true
	start := 0

	for ex || start == 1 {
		ex = false

		rest := workers
		for w := 0; w < workers; w++ {
			offset := w*works + start
			offsetEnd := offset + works - 1

			go func(offset, offsetEnd int) {
				ex := false

				for i := offset; i <= offsetEnd; i += 2 {
					if i+1 < length && array[i] > array[i+1] {
						t := array[i]
						array[i] = array[i+1]
						array[i+1] = t
						ex = true
					}
				}
				ch <- ex
			}(offset, offsetEnd)
		}

		for i := 0; i < rest; i++ {
			if <-ch {
				ex = true
			}
		}

		if start == 0 {
			start = 1
		} else {
			start = 0
		}
	}

	return array
}
