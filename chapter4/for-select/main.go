package main

func main() {

}

func looping() {
	done := make(<-chan string)
	stringStream := make(chan<- string)

	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		}
	}
}

func infinite1() {
	done := make(<-chan string)

	for {
		select {
		case <-done:
			return
		default:
		}
		// do someting
	}
}

func infinite2() {
	done := make(<-chan string)

	for {
		select {
		case <-done:
			return
		default:
			// do someting
		}
	}
}
