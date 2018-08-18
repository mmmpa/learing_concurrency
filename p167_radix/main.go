package p131

const intSize = 32 << (^uint(0) >> 63)

func bit(n int, p uint) int {
	return n >> p & 1
}

func compute(array []int) []int {
	radixSort(array)
	return array
}

type Data struct {
	Head int
	Tail int
	Bit  uint
}

func radixSort(array []int) {
	queue := []Data{{0, len(array) - 1, intSize + 1}}

	for len(queue) != 0 {
		d := queue[0]
		head := d.Head
		tail := d.Tail
		bp := d.Bit

		next := queue[1:]

		if head < tail && bp >= 0 {
			q := part(array, head, tail, bp)
			if bp > 0 {
				next = append(next, Data{head, q, bp - 1}, Data{q + 1, tail, bp - 1})
			}
		}
		queue = next
	}
}

func part(array []int, head, tail int, bp uint) int {
	k := head - 1
	l := tail + 1

	for k++; k < tail && bit(array[k], bp) == 0; k++ {
	}
	for l--; head <= l && bit(array[l], bp) == 1; l-- {
	}

	for k < l {
		t := array[k]
		array[k] = array[l]
		array[l] = t

		for k++; bit(array[k], bp) == 0; k++ {
		}
		for l--; bit(array[l], bp) == 1; l-- {
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
