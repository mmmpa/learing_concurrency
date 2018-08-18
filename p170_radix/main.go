package p131

const intSize = 32 << (^uint(0) >> 63)

func bit(n int, p uint, r uint) int {
	return n >> p & (2<<(r-1) - 1)
}
func compute(array []int) []int {
	return radixSort(array, 3)
}

type Data struct {
	Head int
	Tail int
	Bit  uint
}

func radixSort(array []int, r uint) [] int{
	length := len(array)
	steps := intSize / r

	arrayA := array
	arrayB := make([]int, length)

	countLength := 2<<(r-1)
	
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

		for i := length-1; i >= 0; i-- {
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
	length := len(array)
	queue := make(chan *Data, length)
	done := make(chan int, length)
	rest := length

	for w := 0; w < workers; w++ {
		go func() {
			for d := range queue {
				head := d.Head
				tail := d.Tail
				bp := d.Bit

				if head < tail {
					q := part(array, head, tail, bp)
					if bp == 0 {
						done <- tail - head + 1
					} else {
						queue <- &Data{head, q, bp - 1}
						queue <- &Data{q + 1, tail, bp - 1}
					}
				} else if head == tail {
					done <- 1
				}
			}
		}()
	}

	queue <- &Data{0, len(array) - 1, intSize + 1}

	for rest != 0 {
		rest -= <-done
	}

	return array
}
