package main

import (
	"fmt"
	"sync"
	"time"
)

//WaitGroup必须通过指针传递给函数
func worker(id int, wg *sync.WaitGroup) {
	//return 时，通知 WaitGroup，当前协程的工作已经完成
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(5 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i+1, &wg)
	}
	//主线程阻塞，等待计数器归零
	wg.Wait()
}
