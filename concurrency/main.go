package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true // signal that the greeting is done
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done <- true // signal that the greeting is done
	close(done)  // close the channel to signal no more messages will be sent (only works if we know which goroutine is last)
}

func main() {
	//dones := make([]chan bool, 4)
	done := make(chan bool)

	//dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	//dones[1] = make(chan bool)
	go greet("How are you?", done)
	//dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	//dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	//for _, done := range dones {
	//	<-done // wait for each greeting to finish
	//}

	for range done {
	}
}
