package p131

import (
	"sync"
	"runtime/debug"
)

type Q struct {
	Head   *Item
	Tail   *Item
	Length int
}

func (o *Q) Push(item *Item) {
	o.Length++
	// fmt.Println("increment", o.Length)

	item.Next = nil

	switch {
	case o.Tail == nil:
		o.Head = item
		o.Tail = item
	case o.Tail != nil:
		o.Tail.Next = item
		item.Prev = o.Tail
		o.Tail = item
	}
}

func (o *Q) Pop() *Item {

	switch {
	case o.Tail != nil:
		o.Length--
		// fmt.Println("decrement", o.Length)

		n := o.Tail

		if o.Tail.Prev != nil {
			o.Tail = o.Tail.Prev
		} else {
			o.Head = nil
			o.Tail = nil
		}

		return n
	}
	return nil
}

func (o *Q) Popable() bool {
	return o.Tail != nil
}

type Item struct {
	Head bool
	N    int
	Next *Item
	Prev *Item
	Path *Q
}

func compute(adj [][]int) []*Item {
	nodes := len(adj)
	visited := make([]bool, nodes)
	path := make([]*Item, nodes)
	queue := Q{}

	for i := nodes - 1; i >= 0; i-- {
		item := &Item{N: i, Head: true, Path: &Q{}}
		queue.Push(item)
		path[i] = item
	}

	for {
		node := queue.Pop()

		if node == nil {
			break
		}

		if visited[node.N] {
			continue
		}

		visited[node.N] = true
		node.Path.Push(node)

		for i := nodes - 1; i >= 0; i-- {
			if adj[node.N][i] == 1 {
				queue.Push(&Item{N: i, Path: node.Path})
			}
		}
	}

	return path
}

type Pusher struct {
	I int
	J int
}

func computeC(adj [][]int, workers int) []*Item {
	debug.SetGCPercent(-1)

	nodes := len(adj)
	visited := make([]bool, nodes)
	path := make([]*Item, nodes)
	queue := Q{}
	step := 0

	mu := new(sync.Mutex)
	mu2 := new(sync.Mutex)

	// fmt.Println("start")
	for i := nodes - 1; i >= 0; i-- {
		item := &Item{N: i, Head: true, Path: &Q{}}
		queue.Push(item)
		path[i] = item
	}

	for w := 0; w < workers; w++ {
		go func() {
			for {
				mu2.Lock()
				if queue.Length < 1 {
					continue
				}
				node := queue.Pop()
				mu2.Unlock()

				working := false
				mu.Lock()
				if !visited[node.N] {
					visited[node.N] = true
					node.Path.Push(node)
					step++

					working = true
				}
				mu.Unlock()

				if working {
					for i := nodes - 1; i >= 0; i-- {
						if adj[node.N][i] == 1 {
							mu2.Lock()
							queue.Push(&Item{N: i, Path: node.Path})
							mu2.Unlock()
						}
					}
				}
			}
		}()
	}

	for step < nodes {
	}

	return path
}

func computeA(adj [][]int) []int {
	nodes := len(adj)
	visited := make([]bool, nodes)
	path := make([]int, nodes+nodes)
	step := 0

	for i := 0; i < nodes; i++ {
		path[step] = -1
		step++
		visitA(adj, visited, path, &step, i)
	}

	return path
}

func visitA(adj [][]int, visited []bool, path []int, step *int, node int) {
	if visited[node] {
		return
	}
	nodes := len(adj)

	visited[node] = true
	path[*step] = node
	*step++

	for i := 0; i < nodes; i++ {
		if adj[node][i] == 1 {
			visitA(adj, visited, path, step, i)
		}
	}
}
