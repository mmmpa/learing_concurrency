package p131

import (
	"sync"
	"time"
	"math/rand"
)

type Q struct {
	Head   *Item
	Tail   *Item
	Length int
}

var wait = 10

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
		time.Sleep(time.Duration(wait) * time.Microsecond)

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
	nodes := len(adj)
	visited := make([]bool, nodes)
	path := make([]*Item, nodes)
	queue := Q{}
	mun := workers*2
	finished := false

	poper := make(chan *Item, workers)
	pusher := make(chan *Item, nodes*nodes)
	stepper := make(chan interface{}, nodes)
	closer := make(chan interface{}, nodes)

	lockBox := make([]*sync.Mutex, mun)
	for w := 0; w < mun; w++ {
		lockBox[w] = new(sync.Mutex)
	}

	queueLock := new(sync.Mutex)

	// fmt.Println("start")
	for i := nodes - 1; i >= 0; i-- {
		item := &Item{N: i, Head: true, Path: &Q{}}
		queue.Push(item)
		path[i] = item
	}

	go func() {
		for p := range pusher {
			queueLock.Lock()
			queue.Push(p)
			queueLock.Unlock()
		}
	}()

	go func() {
		defer close(poper)
		for !finished {
			queueLock.Lock()
			p := queue.Pop()
			queueLock.Unlock()
			if p != nil {
				poper <- p
			} else {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Nanosecond)
			}
		}
	}()

	go func() {
		for w := 0; w < workers; w++ {
			<-closer
		}
		close(pusher)
	}()

	for w := 0; w < workers; w++ {
		go func() {
		WORK:
			for node := range poper {
				working := false

				m := muIndex(mun, node.N)
				lockBox[m].Lock()
				if !visited[node.N] {
					visited[node.N] = true
					working = true
				}
				lockBox[m].Unlock()

				if working {
					node.Path.Push(node)
					stepper <- struct{}{}
					time.Sleep(time.Duration(wait) * time.Microsecond)

					for i := nodes - 1; i >= 0; i-- {
						if adj[node.N][i] == 1 {
							if finished {
								break WORK
							}
							pusher <- &Item{N: i, Path: node.Path}
						}
					}
				}
				if finished {
					break WORK
				}
			}
			closer <- struct{}{}
		}()
	}

	for i := 0; i < nodes; i++ {
		<-stepper
	}

	finished = true
	close(stepper)

	return path
}

func muIndex(w, n int) int {
	if n < w {
		return n
	} else {
		return n % w
	}
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
