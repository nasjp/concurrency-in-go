package main

import "fmt"

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case valueStream <- v:
				case <-done:
					return
				}
			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case takeStream <- <-valueStream:
			case <-done:
			}
		}
	}()
	return takeStream
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case valueStream <- fn():
			case <-done:
			}
		}
	}()
	return valueStream

}

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case stringStream <- v.(string):
			case <-done:
			}
		}
	}()
	return stringStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	// for num := range take(done, repeat(done, 1), 10) {
	// 	fmt.Printf("%v ", num)
	// }

	// rand := func() interface{} { return rand.Int() }
	// for num := range take(done, repeatFn(done, rand), 10) {
	// 	fmt.Println(num)
	// }

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}
