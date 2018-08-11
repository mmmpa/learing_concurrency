package main

import (
	"math"
)

func compute(rectCount int) float64 {
	sum := 0.0
	width := 1.0 / float64(rectCount)

	for i := 0; i < rectCount; i++ {
		mid := (float64(i) + 0.5) * width
		height := 4.0 / (1.0 + mid*mid)
		sum += height
	}

	return sum * width
}

func computeC(rectCount, workers int) float64 {
	sum := 0.0
	width := 1.0 / float64(rectCount)

	ch := make(chan float64)
	works := int(math.Ceil(float64(rectCount) / float64(workers)))
	head := 0

	for i := 0; i < workers; i++ {
		tail := head + works
		if tail > rectCount {
			tail = rectCount
		}

		go func(ch chan float64, head, tail int) {
			sum := 0.0

			for i := head; i < tail; i++ {
				mid := (float64(i) + 0.5) * width
				height := 4.0 / (1.0 + mid*mid)
				sum += height
			}

			ch <- sum
		}(ch, head, tail)

		head = tail
	}

	for i := 0; i < workers; i++ {
		sum += <-ch
	}

	return sum * width
}

type Message struct {
	Index  int
	Height float64
}

func computeCC(rectCount, workers int) float64 {
	sum := 0.0
	width := 1.0 / float64(rectCount)

	messages := make([]float64, rectCount)
	ch := make(chan Message)
	works := int(math.Ceil(float64(rectCount) / float64(workers)))
	head := 0

	for i := 0; i < workers; i++ {
		tail := head + works
		if tail > rectCount {
			tail = rectCount
		}

		go func(ch chan Message, head, tail int) {
			for i := head; i < tail; i++ {
				mid := (float64(i) + 0.5) * width
				height := 4.0 / (1.0 + mid*mid)

				ch <- Message{
					Index:  i,
					Height: height,
				}
			}
		}(ch, head, tail)

		head = tail
	}

	for i := 0; i < rectCount; i++ {
		message := <-ch
		messages[message.Index] = message.Height
	}

	for _, height := range messages {
		sum += height
	}

	return sum * width
}
