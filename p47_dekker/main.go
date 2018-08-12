package p47

import (
	"fmt"
	"context"
	"time"
)

var favored = 0
var threadZeroWantToEnter = false
var threadOneWantToEnter = false

func run() {
	ctx, _ := context.WithTimeout(
		context.Background(),
		time.Duration(3)*time.Second,
	)

	threadZeroWantToEnter = false
	threadZero(ctx)
	threadOne(ctx)


loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("break!")
			break loop
		}
	}
	time.Sleep(time.Duration(2 * time.Millisecond))
}

func threadZero(ctx context.Context) {
	go func() {
	loop:
		for {
			threadZeroWantToEnter = true
			for threadOneWantToEnter {
				if favored == 1 {
					threadZeroWantToEnter = false
					fmt.Println("wait !!!!")
					for favored == 1 {}
					threadZeroWantToEnter = true
				}
			}
			criticalSectionZero()
			favored = 1
			threadZeroWantToEnter = false

			otherStuffZero()

			select {
			case <-ctx.Done():
				fmt.Println("break 0")
				break loop
			default:
			}
		}
	}()
}

func threadOne(ctx context.Context) {
	go func() {
	loop:
		for {
			threadOneWantToEnter = true
			for threadZeroWantToEnter {
				if favored == 0 {
					threadOneWantToEnter = false
					fmt.Println("wait !!")
					for favored == 0 {}
					threadOneWantToEnter = true
				}
			}
			criticalSectionOne()
			favored = 0
			threadOneWantToEnter = false

			otherStuffOne()

			select {
			case <-ctx.Done():
				fmt.Println("break 1")
				break loop
			default:
			}
		}
	}()
}

func criticalSectionZero() {
	time.Sleep(time.Duration(1 * time.Millisecond))
	fmt.Println("criticalSection Zero")
}

func otherStuffZero() {
	time.Sleep(time.Duration(1 * time.Millisecond))
	fmt.Println("otherStuff Zero")
}

func criticalSectionOne() {
	time.Sleep(time.Duration(1 * time.Millisecond))
	fmt.Println("criticalSection One")
}

func otherStuffOne() {
	time.Sleep(time.Duration(5 * time.Millisecond))
	fmt.Println("otherStuff One")
}
