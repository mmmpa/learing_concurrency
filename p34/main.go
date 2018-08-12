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
	ch := make(chan SplitGrid)
	works := common.SplitWorks(h, workers)
	head := 1

	for _, n := range works {
		grid := make(Grid, n+2)
		nextGrid := make(Grid, n+2)

		for y := 0; y < len(grid); y++ {
			grid[y] = make([]bool, w+2)
			nextGrid[y] = make([]bool, w+2)

			for x, b := range current[head+y-1] {
				grid[y][x] = b
				nextGrid[y][x] = b
			}
		}

		go func(ch chan SplitGrid, head int, current, next Grid, w, h int) {
			for y := 1; y <= h; y++ {
				for x := 1; x <= w; x++ {
					next[y][x] = detectAlive(current, x, y)
				}
			}
			ch <- SplitGrid{
				Head: head,
				Grid: next,
			}
		}(ch, head, grid, nextGrid, w, n)

		head += n
	}

	for i := 0; i < workers; i++ {
		result := <-ch

		for i, cols := range result.Grid[1:len(result.Grid)-1] {
			next[result.Head+i] = cols
		}
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
