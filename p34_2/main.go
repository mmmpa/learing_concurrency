package p34

import (
	"github.com/mmmpa/parallel/common"
)

type Grid = [][]bool
type SplitGrid struct {
	Head int
	Grid Grid
}

func compute(current, next Grid, w, h int) Grid {
	for y := 1; y <= h; y++ {
		for x := 1; x <= w; x++ {
			next[y][x] = detectAlive(current, x, y)
		}
	}

	return next
}

func computeC(current, next Grid, w, h, workers int) Grid {
	ch := make(chan interface{})
	defer close(ch)

	works := common.SplitWorks(h, workers)
	head := 1

	for _, n := range works {
		go func(head, n int) {
			for y := head; y < head+n; y++ {
				for x := 1; x <= w; x++ {
					next[y][x] = detectAlive(current, x, y)
				}
			}
			ch <- struct{}{}
		}(head, n)

		head += n
	}

	for i := 0; i < workers; i++ {
		<-ch
	}

	return next
}

func detectAlive(grid Grid, x, y int) bool {
	count := 0
	if grid[y-1][x-1] {
		count++
	}
	if grid[y-1][x] {
		count++
	}
	if grid[y-1][x+1] {
		count++
	}

	if grid[y][x-1] {
		count++
	}
	if grid[y][x+1] {
		count++
	}

	if grid[y+1][x-1] {
		count++
	}
	if grid[y+1][x] {
		count++
	}
	if grid[y+1][x+1] {
		count++
	}

	switch {
	case count == 1:
		return false
	case count == 2 && grid[y][x]:
		return true
	case count == 3:
		return true
	default:
		return false
	}
}
