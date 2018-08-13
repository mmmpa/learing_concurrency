package p101

import (
	"sort"
)

const (
	Q = 5

	LESS    = 0
	MED     = 1
	GREATER = 2
)

func compute(array []int, index int) int {
	l := len(array)

	if l <= Q {
		return pick(array, index)
	}

	split := l/Q + 1
	medians := make([]int, split)

	for i := 0; i < split; i++ {
		splitMedian := i*(Q-1) + Q/2

		if splitMedian > l-1 {
			medians[i] = array[l-1]
		} else {
			medians[i] = array[splitMedian]
		}
	}

	median := compute(medians, split/2)

	counts := make([]int, 3)
	marks := make([]int, l)
	countAndMark(array, median, counts, marks)

	k := index + 1

	switch {
	case k <= counts[LESS]:
		next := make([]int, counts[LESS])
		pack(array, marks, next, LESS)
		return compute(next, index)
	case counts[LESS]+counts[MED] < k:
		next := make([]int, counts[GREATER])
		pack(array, marks, next, GREATER)
		return compute(next, index-(counts[LESS]+counts[MED]))
	default:
		return median
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
