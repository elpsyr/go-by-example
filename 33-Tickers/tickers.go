package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		i := 0
		for {
			i += 1
			fmt.Println(i)
			select {
			case <-done:
				fmt.Println("done")
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true

	time.Sleep(2 * time.Second)
	fmt.Println("Ticker stopped")

}
