package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		//进门
		wg.Add(1)
		go func() {
			fmt.Println("I'm ", time.Now())

			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				//ops+=1
			}
			//出门
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}
