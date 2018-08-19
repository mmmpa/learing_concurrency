package p131

import (
	"context"
)

const intSize = 32 << (^uint(0) >> 63)

func compute(array []int, target int) int {
	for i, n := range array {
		if n == target {
			return i
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
