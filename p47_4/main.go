package p47

import (
	"fmt"
	"context"
	"time"
	"math/rand"
)

var threadZeroWantToEnter = true
var threadOneWantToEnter = true

func randomSleep(){
	time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
}

func run() {
	ctx, _ := context.WithTimeout(
		context.Background(),
		time.Duration(3)*time.Second,
	)
	ch := make(chan interface{})

	threadZeroWantToEnter = false
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
			threadZeroWantToEnter = true
			for threadOneWantToEnter {
				threadZeroWantToEnter = false
				randomSleep()
				threadZeroWantToEnter = true
			}
			criticalSectionZero()
			threadZeroWantToEnter = false
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
			threadOneWantToEnter = true
			for threadZeroWantToEnter {
				threadOneWantToEnter = false
				randomSleep()
				threadOneWantToEnter = true
			}
			criticalSectionOne()
			threadOneWantToEnter = false

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
	ctx, _ := context.WithTimeout(
		context.Background(),
		time.Duration(100)*time.Millisecond,
	)
	<-ctx.Done()
	fmt.Println("otherStuff Zero")
}

func criticalSectionOne() {
	fmt.Println("criticalSection One")
}

func otherStuffOne() {
	ctx, _ := context.WithTimeout(
		context.Background(),
		time.Duration(500)*time.Millisecond,
	)
	<-ctx.Done()
	fmt.Println("otherStuff One")
}
