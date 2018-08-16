package p131

func compute(array []int) []int {
	length := len(array)

	h := 1
	for h < length {
		h = 3*h + 1
	}
	h /= 3

	for h != 1 {
		for k := 0; k < h; k++ {
			for i := k; i < length; i += h {
				v := array[i]
				j := i

				for j-h >= 0 && array[j-h] > v {
					array[j] = array[j-h]

					j -= h
					if j <= h {
						break
					}
				}
				array[j] = v
			}
		}

		h /= 3
	}

	return insert(array)
}

func insert(array []int) []int {
	length := len(array)

	for i := 1; i < length; i++ {
		v := array[i]
		j := i
		for array[j-1] > v {
			array[j] = array[j-1]

			j--
			if j <= 0 {
				break
			}
		}
		array[j] = v
	}

	return array
}

func computeC(array []int, workers int) []int {
	length := len(array)
	ch := make(chan interface{})

	h := 1
	for h < length {
		h = 3*h + 1
	}
	h /= 3

	for h != 1 {
		works := h / workers
		if h%workers != 0 {
			works++
		}

		for w := 0; w < workers; w++ {
			head := w * works
			tail := head + works - 1
			if tail > h {
				tail = h - 1
			}

			go func(head, tail int) {
				for k := head+h; k <= tail; k++ {
					for i := k; i < length; i += h {
						v := array[i]
						j := i

						for array[j-h] > v {
							array[j] = array[j-h]

							j -= h
							if j <= h {
								break
							}
						}
						array[j] = v
					}
				}

				ch <- struct{}{}
			}(head, tail)
		}

		for i := 0; i < workers; i++ {
			<-ch
		}

		h /= 3
	}

	return insert(array)
}
