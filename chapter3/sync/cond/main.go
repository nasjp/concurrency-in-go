package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// queue()
	button()
}

func queue() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(deley time.Duration) {
		time.Sleep(deley)
		c.L.Lock()
		queue = queue[1:]
		fmt.Printf("Removed from queue: %d\n", len(queue))
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		queue = append(queue, struct{}{})
		fmt.Printf("Added to queue: %d\n", len(queue))
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

func button() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegisterd sync.WaitGroup
	clickRegisterd.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegisterd.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegisterd.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegisterd.Done()
	})

	button.Clicked.Broadcast()
	clickRegisterd.Wait()
}
