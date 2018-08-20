package p131

import (
	"context"
)

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
func computeC(array []int, target int, workers int) int {
	length := len(array)
	result := -1
	not := make(chan interface{})

	for w := 0; w < workers; w++ {
		go func(offset int) {
			for i := offset; i < length; i += workers {
				if array[i] == target {
					result = i
					return
				} else if result != -1 {
					return
				}
			}
			not <- struct{}{}
		}(w)
	}

	w := 0
	for {
		select {
		case <-not:
			w++
			if w == workers {
				return -1
			}
		default:
			if result != -1 {
				return result
			}
		}
	}
}

func computeCC(array []int, target int, workers int) int {
	length := len(array)
	finish := make(chan int)
	not := make(chan interface{})

	ctx, cancel := context.WithCancel(context.Background())

	for w := 0; w < workers; w++ {
		go func(offset int) {
			for i := offset; i < length; i += workers {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if array[i] == target {
					finish <- i
					break
				}
			}
			not <- struct{}{}
		}(w)
	}

	for w := 0; w < workers; w++ {
		select {
		case i := <-finish:
			cancel()
			return i
		case <-not:
		}
	}

	return -1
}
