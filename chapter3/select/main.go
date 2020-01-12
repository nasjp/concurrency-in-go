package main

import (
	"fmt"
	"time"
)

func main() {
	// halfSelect()
	// timeout()
	// defaultGate()
	otherWork()
}

func simpleSelect() {
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func halfSelect() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	close(c1)
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}

	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func timeout() {
	var c <-chan int

	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("time out.")
	}
}

func defaultGate() {
	start := time.Now()
	var c1, c2 <-chan int

	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n", time.Since(start))
	}
}

func otherWork() {
	done := make(chan interface{})
	go func() {
		defer close(done)
		time.Sleep(5 * time.Second)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
