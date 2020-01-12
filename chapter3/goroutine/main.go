package main

import (
	"fmt"
	"sync"
)

func main() {
	// helloWorld()
	// welcome()
	helloGreetingsGoodDay()
}

func helloWorld() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello world")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()
}

func welcome() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()

	wg.Wait()
	fmt.Println(salutation)
}

func helloGreetingsGoodDay() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
