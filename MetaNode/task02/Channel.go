package main

import (
	"fmt"
	"sync"
)

/**
 * 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
 * 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
 */
func send(ch chan int, group *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("发送：", i)
	}
	close(ch)
	defer group.Done()
}
func receive(ch chan int, group *sync.WaitGroup) {
	for v := range ch {
		fmt.Println("接收：", v)
	}
	defer group.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int, 10)
	go send(ch, &wg)
	go receive(ch, &wg)
	wg.Wait()
}
