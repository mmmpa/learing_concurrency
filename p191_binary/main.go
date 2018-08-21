package p131

const intSize = 32 << (^uint(0) >> 63)

func compute(array []int, target int, n int) int {
	length := len(array)

	mid := make([]int, n+1)
	lr := make([]string, n+2)

	head := 0
	tail := length - 1
	result := -1

	lr[0] = "R"
	lr[n+1] = "L"
	for head <= tail && result == -1 {
		mid[0] = head - 1
		step := float64(tail-head+1) / float64(n+1)

		for i := 1; i <= n; i++ {
			offset := int(step*float64(i) + float64(i-1))
			mid[i] = head + offset

			if mid[i] <= tail {
				v := array[mid[i]]
				switch {
				case target < v:
					lr[i] = "L"
				case v < target:
					lr[i] = "R"
				default:
					lr[i] = "E"
					result = mid[i]
				}
			} else {
				mid[i] = tail + 1
				lr[i] = "L"
			}
		}

		for i := 1; i <= n; i++ {
			if lr[i-1] != lr[i] {
				head = mid[i-1] + 1
				tail = mid[i] - 1
			}
		}

		// 全て R の時
		if lr[n] != lr[n+1] {
			head = mid[n] + 1
		}
	}

	return result
}

func computeA(array []int, target int) int {
	length := len(array)

	head := 0
	tail := length - 1

	for head <= tail {
		mid := head + (tail-head)/2
		v := array[mid]

		switch {
		case target < v:
			tail = mid - 1
		case v < target:
			head = mid + 1
		default:
			return target
		}
	}

	return -1
}

type Data struct {
	Step float64
	Head int
	I    int
}

func farm(array []int, target int, n int, mid []int, lr []string, tail int, producer chan Data, result chan int) chan interface{} {
	not := make(chan interface{})

	for i := 1; i <= n; i++ {
		go func(i int) {
			for data := range producer {
				head := data.Head
				i := data.I
				step := data.Step

				offset := int(step*float64(i) + float64(i-1))
				mid[i] = head + offset

				if mid[i] <= tail {
					v := array[mid[i]]
					switch {
					case target < v:
						lr[i] = "L"
					case v < target:
						lr[i] = "R"
					default:
						lr[i] = "E"
						result <- mid[i]
					}
				} else {
					mid[i] = tail + 1
					lr[i] = "L"
				}
				not <- struct{}{}
			}
		}(i)
	}

	return not
}

func computeC(array []int, target int, n int) int {
	length := len(array)

	mid := make([]int, n+1)
	lr := make([]string, n+2)

	head := 0
	tail := length - 1

	lr[0] = "R"
	lr[n+1] = "L"

	producer := make(chan Data)
	result := make(chan int)
	not := farm(array, target, n, mid, lr, tail, producer, result)

	for head <= tail {
		mid[0] = head - 1
		step := float64(tail-head+1) / float64(n+1)

		for i := 1; i <= n; i++ {
			producer <- Data{
				Head: head,
				Step: step,
				I:    i,
			}
		}

		w := 0
	WAIT:
		for {
			select {
			case <-not:
				w++
				if w == n {
					break WAIT
				}
			case re := <-result:
				return re
			default:
			}
		}

		for i := 1; i <= n; i++ {
			if lr[i-1] != lr[i] {
				head = mid[i-1] + 1
				tail = mid[i] - 1
			}
		}

		// 全て R の時
		if lr[n] != lr[n+1] {
			head = mid[n] + 1
		}
	}

	return -1
}
