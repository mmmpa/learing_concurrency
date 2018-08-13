package p101

import (
	"sort"
	"github.com/mmmpa/parallel/p93_prefix_scan"
)

const (
	Q = 5
	M = 3

	LESS    = 0
	MED     = 1
	GREATER = 2
)

func compute(array []int, index int) int {
	l := len(array)

	if l <= Q {
		return pick(array, index)
	}

	cellsCount := l/Q + 1
	medians := make([]int, cellsCount)

	for i := 0; i < cellsCount; i++ {
		splitMedian := i*(Q-1) + M

		if splitMedian > l-1 {
			medians[i] = array[l-1]
		} else {
			medians[i] = array[splitMedian]
		}
	}

	median := compute(medians, cellsCount/2)

	counts := make([]int, 3)
	marks := make([]int, l)
	countAndMark(array, median, counts, marks)

	k := index + 1
	lessEqualIndex := counts[LESS] + counts[MED]

	switch {
	case k <= counts[LESS]:
		next := make([]int, counts[LESS])
		pack(array, marks, next, LESS)
		return compute(next, index)
	case k <= lessEqualIndex:
		return median
	default:
		next := make([]int, counts[GREATER])
		pack(array, marks, next, GREATER)
		return compute(next, index-lessEqualIndex)
	}
}

func pick(array []int, index int) int {
	sort.Ints(array)
	return array[index]
}

func countAndMark(array []int, m int, counts []int, marks []int) {
	for i, n := range array {
		switch {
		case n < m:
			marks[i] = LESS
			counts[LESS]++
		case m == n:
			marks[i] = MED
			counts[MED]++
		case m < n:
			marks[i] = GREATER
			counts[GREATER]++
		}
	}
}

func pack(array []int, marks []int, container []int, key int) {
	j := 0
	for i, mark := range marks {
		if mark == key {
			container[j] = array[i]
			j++
		}
	}
}

func computeC(array []int, index int, workers int) int {
	l := len(array)

	if l <= Q {
		return pick(array, index)
	}

	split := l/Q + 1
	medians := findMedians(array, workers)

	median := computeC(medians, split/2, workers)

	counts := make([]int, 3)
	marks := make([]int, l)
	countAndMarkC(array, median, counts, marks, workers)

	k := index + 1
	lessEqualIndex := counts[LESS] + counts[MED]

	switch {
	case k <= counts[LESS]:
		next := make([]int, counts[LESS])
		packC(array, marks, next, LESS, workers)
		return compute(next, index)
	case k <= lessEqualIndex:
		return median
	default:
		next := make([]int, counts[GREATER])
		packC(array, marks, next, GREATER, workers)
		return compute(next, index-lessEqualIndex)
	}
}

func packC(array []int, marks []int, container []int, key int, workers int) {
	l := len(array)
	ch := make(chan interface{})
	flags := make([]int, len(marks))

	for i := 0; i < workers; i++ {
		go func(offset int) {
			for i := offset; i < l; i += workers {
				if marks[i] == key {
					flags[i] = 1
				}
			}
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < workers; i++ {
		<-ch
	}

	scans := p93.PrefixScan(flags, workers)
	container[0] = scans[0]

	for i := 0; i < workers; i++ {
		go func(offset int) {
			for i := offset + 1; i < l; i += workers {
				if scans[i] != scans[i-1] {
					container[scans[i]-1] = array[i]
				}
			}
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < workers; i++ {
		<-ch
	}
}

func countAndMarkC(array []int, m int, counts []int, marks []int, workers int) {
	l := len(array)
	ch := make(chan []int)

	for i := 0; i < workers; i++ {
		go func(offset int) {
			counts := make([]int, 3)

			for i := offset; i < l; i += workers {
				n := array[i]
				switch {
				case n < m:
					marks[i] = LESS
					counts[LESS]++
				case m == n:
					marks[i] = MED
					counts[MED]++
				case m < n:
					marks[i] = GREATER
					counts[GREATER]++
				}
			}
			ch <- counts
		}(i)
	}

	for i := 0; i < workers; i++ {
		partCounts := <-ch
		counts[LESS] += partCounts[LESS]
		counts[MED] += partCounts[MED]
		counts[GREATER] += partCounts[GREATER]
	}
}

func findMedians(array []int, defaultWorkers int) []int {
	l := len(array)
	ch := make(chan []int)
	workers := defaultWorkers

	if Q*defaultWorkers > l {
		workers = l/Q + 1
	}

	ll := l / workers
	if l%workers != 0 {
		ll++
	}

	totalCells := 0
	for i := 0; i < workers; i++ {
		head := i * ll
		tail := head + ll

		if tail > l-1 {
			tail = l - 1
		}

		length := tail - head
		cellsCount := length/Q + 1
		totalCells += cellsCount

		go func(head, tail, cellsCount int) {
			medians := make([]int, cellsCount)

			for i := 0; i < cellsCount; i++ {
				cellMedian := head + i*(Q-1) + M

				if cellMedian > tail {
					medians[i] = array[tail]
				} else {
					medians[i] = array[cellMedian]
				}
			}

			ch <- medians
		}(head, tail, cellsCount)
	}

	medians := make([]int, totalCells)

	j := 0
	for i := 0; i < workers; i++ {
		for _, n := range <-ch {
			medians[j] = n
			j++
		}
	}

	return medians
}
