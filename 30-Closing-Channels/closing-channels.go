package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 100; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	//关闭通道，告诉接收方接收工作已完成
	close(jobs)

	//主线程在等待两个协程执行完毕
	fmt.Println("sent all jobs")
	fmt.Println(<-done)
	//<-done
}
