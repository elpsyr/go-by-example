package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	//一个非空的通道也可以被关闭，通道中的值依旧能取到
	//管子变成了桶，见底了就知道没了
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
