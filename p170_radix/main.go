package p131

const intSize = 32 << (^uint(0) >> 63)

func bit(n int, p uint, r uint) int {
	return n >> p & (2<<(r-1) - 1)
}
func compute(array []int) []int {
	return radixSortA(array, 3)
}

type Data struct {
	Head int
	Tail int
	Bit  uint
}

func radixSortC(array []int, r uint, workers int) [] int {
	length := len(array)
	steps := intSize / r

	done := make(chan interface{}, workers)

	arrayA := array
	arrayB := make([]int, length)

	countLength := 2 << (r - 1)

	workerRange := length / workers
	if length%workers != 0 {
		workerRange++
	}

	localStore := make([][]int, workers)

	for ii := uint(0); ii < steps; ii++ {
		for i, _ := range localStore {
			localStore[i] = make([]int, countLength)
		}

		head := 0
		for w := 0; w < workers; w++ {
			tail := head + workerRange + 1
			if tail > length {
				tail = length
			}

			go func(id, head, tail int) {
				for i := head; i < tail; i++ {
					v := arrayA[i]
					bt := bit(v, r*ii, r)

					localStore[id][bt]++
				}
				done <- struct{}{}
			}(w, head, tail)

			head = tail
		}

		for w := 0; w < workers; w++ {
			<-done
		}

		count := make([]int, countLength)
		count2 := make([]int, countLength)

		for i := 0; i < countLength; i++ {
			for _, ls := range localStore {
				count[i] += ls[i]
			}
		}

		count2[0] = -1
		count2[1] = count[0] - 1
		for i := 2; i < countLength; i++ {
			count2[i] += count2[i-1] + count[i-1]
		}

		for _, ls := range localStore {
			for i := 0; i < countLength; i++ {
				n := ls[i]
				ls[i] += count2[i]
				count2[i] += n
			}
		}

		head = 0
		for w := 0; w < workers; w++ {
			tail := head + workerRange + 1
			if tail > length {
				tail = length
			}

			go func(id, head, tail int) {
				for i := tail - 1; i >= head; i-- {
					v := arrayA[i]
					bt := bit(v, r*ii, r)

					p := localStore[id][bt]
					localStore[id][bt]--
					arrayB[p] = v
				}
				done <- struct{}{}
			}(w, head, tail)

			head = tail
		}

		for w := 0; w < workers; w++ {
			<-done
		}

		t := arrayB
		arrayB = arrayA
		arrayA = t
	}

	return arrayA
}

func radixSort(array []int, r uint) [] int {
	length := len(array)
	steps := intSize / r

	arrayA := array
	arrayB := make([]int, length)

	countLength := 2 << (r - 1)

	for ii := uint(0); ii < steps; ii++ {
		offset := -1

		for countPos := 0; countPos < countLength; countPos++ {
			count := 0

			for i := 0; i < length; i++ {
				v := arrayA[i]
				bt := bit(v, r*ii, r)

				if bt == countPos {
					count++
				}
			}

			rank := offset + count

			for i := length - 1; i >= 0; i-- {
				v := arrayA[i]
				bt := bit(v, r*ii, r)

				if bt == countPos {
					arrayB[rank] = v
					rank--
				}
			}

			offset += count
		}

		t := arrayB
		arrayB = arrayA
		arrayA = t
	}

	return arrayA
}

func radixSortA(array []int, r uint) [] int {
	length := len(array)
	steps := intSize / r

	arrayA := array
	arrayB := make([]int, length)

	countLength := 2 << (r - 1)

	for ii := uint(0); ii < steps; ii++ {
		count := make([]int, countLength)

		for i := 0; i < length; i++ {
			v := arrayA[i]
			bt := bit(v, r*ii, r)
			count[bt]++
		}

		count[0]--

		for i := 1; i < countLength; i++ {
			count[i] = count[i] + count[i-1]
		}

		for i := length - 1; i >= 0; i-- {
			v := arrayA[i]
			bt := bit(v, r*ii, r)

			arrayB[count[bt]] = v
			count[bt]--
		}

		t := arrayB
		arrayB = arrayA
		arrayA = t
	}

	return arrayA
}

func part(array []int, head, tail int, bp uint) int {
	k := head - 1
	l := tail + 1

	for k++; k < tail && bit(array[k], bp, 1) == 0; k++ {
	}
	for l--; head <= l && bit(array[l], bp, 1) == 1; l-- {
	}

	for k < l {
		t := array[k]
		array[k] = array[l]
		array[l] = t

		for k++; bit(array[k], bp, 1) == 0; k++ {
		}
		for l--; bit(array[l], bp, 1) == 1; l-- {
		}
	}

	return l
}

func computeC(array []int, workers int) []int {
	return radixSortC(array, 3, workers)
}
