package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()

	go func(in chan string) {
		fmt.Println(in)

		fmt.Println(<-in)
	}(messages)

	msg := <-messages
	fmt.Println(msg)
}
