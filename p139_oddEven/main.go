package p131

import (
	"time"
	"math/rand"
)

func compute(array []int) []int {
	length := len(array)

	ex := true
	start := 0

	for ex || start == 1 {
		ex = false

		for i := start; i < length-1; i += 2 {
			if array[i] > array[i+1] {
				t := array[i]
				array[i] = array[i+1]
				array[i+1] = t
				ex = true
			}
		}

		if start == 0 {
			start = 1
		} else {
			start = 0
		}
	}

	return array
}

func computeC(array []int, workers int) []int {
	length := len(array)
	ch := make(chan interface{})

	works := length/workers + 1
	if works%2 != 0 {
		works++
	}

	ex := true
	start := 0

	for ex || start == 1 {
		ex = false

		rest := workers
		for w := 0; w < workers; w++ {
			offset := w*works + start
			offsetEnd := offset + works - 1

			go func(offset, offsetEnd int) {
				for i := offset; i <= offsetEnd; i += 2 {
					if i+1 > length-1 {
						break
					}

					if array[i] > array[i+1] {
						t := array[i]
						array[i] = array[i+1]
						array[i+1] = t
						ex = true
					}
				}

				ch <- struct{}{}
			}(offset, offsetEnd)
		}

		for i := 0; i < rest; i++ {
			<-ch
		}

		if start == 0 {
			start = 1
		} else {
			start = 0
		}
	}

	return array
}

func computeCC(array []int, workers int) []int {
	length := len(array)
	ch0 := make(chan int)
	ch1 := make(chan bool)

	works := length/workers + 1
	if works%2 != 0 {
		works++
	}

	zero := 0
	one := 0
	trip := false

	for w := 0; w < workers; w++ {
		offset := w * works
		offsetEnd := offset + works - 1

		go func(offset, offsetEnd, id int) {
			for {
				ex0 := false
				ex1 := false

			zero:
				for i := offset; i <= offsetEnd; i += 2 {
					if i+1 > length-1 {
						break zero
					}

					if array[i] > array[i+1] {
						t := array[i]
						array[i] = array[i+1]
						array[i+1] = t
						ex0 = true
					}
				}

				ch0 <- 1

				for zero < workers {
					time.Sleep(time.Duration(rand.Intn(1)) * time.Nanosecond)
				}

			one:
				for i := offset + 1; i <= offsetEnd+1; i += 2 {
					if i+1 > length-1 {
						break one
					}

					if array[i] > array[i+1] {
						t := array[i]
						array[i] = array[i+1]
						array[i+1] = t
						ex1 = true
					}
				}

				ch1 <- ex0 || ex1

				for one < workers {
					time.Sleep(time.Duration(rand.Intn(1)) * time.Nanosecond)
				}

				trip = true
			}
		}(offset, offsetEnd, w)
	}

	go func() {
		for {
			zero += <-ch0
			if zero >= workers {
				one = 0
			}
		}
	}()

	ex := true

	go func() {
		lex := false
		for {
			if <-ch1 {
				lex = true
			}
			one ++
			if one >= workers {
				if !lex {
					ex = false
				}
				zero = 0
				lex = false
			}
		}
	}()

	for !trip || ex {

	}

	return array
}
