package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func main() {
	t := time.Now()

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	//创建三个协程去干活
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	//派发工作
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	//再次派发工作
	//panic: send on closed channel
	//for j := 1; j <= numJobs; j++ {
	//	jobs<-j*2
	//}

	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}

	elapsed := time.Since(t)

	fmt.Println("app elapsed:", elapsed)
}
