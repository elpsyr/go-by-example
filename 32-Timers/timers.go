package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	//主线程阻塞2s
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(100 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	//主线程阻塞5s
	time.Sleep(5 * time.Second)

}
