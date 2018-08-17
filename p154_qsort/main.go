package p131

import (
	"sync"
	"fmt"
	"time"
)

func compute(array []int) []int {
	qSort(array)
	return array
}

func computeQ(array []int) []int {
	qSort2(array, 0, len(array)-1)
	return array
}

type Data struct {
	Head int
	Tail int
}

func qSort2(array []int, head, tail int) {
	if head < tail {
		q := part(array, head, tail)
		qSort2(array, head, q-1)
		qSort2(array, q+1, tail)
	}
}

func qSort(array []int) {
	queue := []Data{{0, len(array) - 1}}

	for len(queue) != 0 {
		var next []Data
		for _, data := range queue {
			head := data.Head
			tail := data.Tail

			if head < tail {
				q := part(array, head, tail)
				next = append(next, Data{head, q - 1})
				next = append(next, Data{q + 1, tail})
			}
		}
		queue = next
	}
}

func part(array []int, head, tail int) int {
	x := array[head]
	k := head
	l := tail + 1

	for k++; array[k] <= x && k < tail; k++ {
	}
	for l--; array[l] > x; l-- {
	}

	for k < l {
		t := array[k]
		array[k] = array[l]
		array[l] = t

		for k++; array[k] <= x; k++ {
		}
		for l--; array[l] > x; l-- {
		}
	}
	t := array[head]
	array[head] = array[l]
	array[l] = t

	return l
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
	qSortC(array, workers)
	return array
}

func qSortC(array []int, workers int) {
	mu := new(sync.RWMutex)
	mu2 := new(sync.RWMutex)
	queue := []Data{{0, len(array) - 1}}
	rest := 1

	enqueue := make(chan []Data)
	dequeue := make(chan Data)

	go func() {
		for {
			data := <-enqueue
			mu.Lock()
			queue = append(queue, data...)
			mu.Unlock()
		}
	}()

	go func() {
		for {
			for len(queue) == 0 {
				time.Sleep(1 * time.Nanosecond)
			}

			mu.Lock()
			l := len(queue)
			d := queue[0]
			if l == 1 {
				queue = []Data{}
			} else {
				queue = queue[1:l]
			}
			mu.Unlock()
			dequeue <- d
		}
	}()

	for w := 0; w < workers; w++ {
		go func() {
			for {
				data := <-dequeue

				head := data.Head
				tail := data.Tail

				r := -1
				if head < tail {
					q := part(array, head, tail)
					enqueue <- []Data{{head, q - 1}, {q + 1, tail}}
					r += 2
				}

				mu2.Lock()
				rest += r
				mu2.Unlock()
			}
		}()
	}

	for rest != 0 {
	}

	fmt.Println("end")
}
