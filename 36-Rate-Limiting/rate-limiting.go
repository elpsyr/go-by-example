package main

import (
	"fmt"
	"time"
)

func main() {
	//发布5个请求
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i + 1
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	//每200ms公布一个请求
	for req := range requests {

		<-limiter //释放一个时间点，敲一下钟
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 4)

	//放三个时间名额进去
	for i := 0; i < 4; i++ {
		burstyLimiter <- time.Now()
	}
	//每隔200ms放一个时间名额进去
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//5个请求放入通道
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("solve request", req, time.Now())

	}

}
