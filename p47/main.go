package p47

import (
	"fmt"
	"context"
	"time"
)

var threadNumber = 0

func run() {
	ctx, _ := context.WithTimeout(
		context.Background(),
		time.Duration(3)*time.Second,
	)
	ch := make(chan interface{})

	threadZero(ch)
	threadOne(ch)

loop:
	for {
		select {
		case <-ctx.Done():
			ch <- struct{}{}
			ch <- struct{}{}
			fmt.Println("break!")
			break loop
		}
	}
}

func threadZero(ch chan interface{}) {
	go func() {
	loop:
		for {
			for threadNumber == 1 {
				// スピンウェイト
			}
			criticalSectionZero()
			threadNumber = 1
			otherStuffZero()

			select {
			case <-ch:
				fmt.Println("break 0")
				break loop
			default:
			}
		}
	}()
}

func threadOne(ch chan interface{}) {
	go func() {
	loop:
		for {
			for threadNumber == 0 {
				// スピンウェイト
			}
			criticalSectionOne()
			threadNumber = 0
			otherStuffOne()

			select {
			case <-ch:
				fmt.Println("break 1")
				break loop
			default:
			}
		}
	}()
}

func criticalSectionZero() {
	fmt.Println("criticalSection Zero")
}

func otherStuffZero() {
	fmt.Println("otherStuff Zero")
}
func criticalSectionOne() {
	fmt.Println("criticalSection One")
}

func otherStuffOne() {
	fmt.Println("otherStuff One")
}
