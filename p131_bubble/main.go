package p131

import (
	"time"
	"math/rand"
)

func compute(array []int) []int {
	length := len(array)

	for i := length - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				t := array[j]
				array[j] = array[j+1]
				array[j+1] = t
			}
		}
	}

	return array
}

func computeC(array []int, rowWorkers int) []int {
	ch := make(chan interface{})
	length := len(array)

	workers := rowWorkers
	if workers > length-1 {
		workers = length - 1
	}

	worked := make([]int, workers)
	var finished interface{}

	for id := 0; id < workers; id++ {
		pre := id - 1
		if pre < 0 {
			pre = workers - 1
		}
		next := id + 1
		if next > workers-1 {
			next = 0
		}

		allEnd := length - 2
		go func(ch chan interface{}, id, pre, next int) {
			r := 0
			for {
				end := allEnd - id - workers*r

				if end < 0 {
					break
				}

				for j := 0; j <= end; j++ {
					for !(r == 0 && id == 0) && worked[pre] <= j {
						time.Sleep(time.Duration(rand.Intn(1)) * time.Microsecond)
					}
					if array[j] > array[j+1] {
						t := array[j]
						array[j] = array[j+1]
						array[j+1] = t
					}

					worked[id] = j
				}

				if end == 0 {
					ch <- struct{}{}
					break
				}

				//  // 手前のワーカーが-次のターンに入っていない
				//                              // 次のワーカーが追いついていない
				for worked[pre] > worked[id] || worked[id]-1 > worked[next] {
					time.Sleep(time.Duration(rand.Intn(1)) * time.Microsecond)

					if finished != nil {
						break
					}
				}

				r++
			}
		}(ch, id, pre, next)
	}

	finished = <-ch

	return array
}
