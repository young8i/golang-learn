package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
 */
func printOddNumbers(group *sync.WaitGroup) {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数：", i)
	}
	defer group.Done()
}
func printEvenNumbers(group *sync.WaitGroup) {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数：", i)
	}
	defer group.Done()
}

/**
 * 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
 */

func TaskScheduler(tasks []func()) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func() {
			defer wg.Done()
			startTime := time.Now()
			task()
			endTime := time.Now()
			executionTime := endTime.Sub(startTime)
			fmt.Printf("任务执行时间：%v\n", executionTime)
		}()
	}
	wg.Wait()
}

func main() {
	fmt.Println("==================奇数偶数任务开始执行==================")
	var wg sync.WaitGroup
	wg.Add(2)
	go printOddNumbers(&wg)
	go printEvenNumbers(&wg)
	wg.Wait()
	fmt.Println("==================奇数偶数任务执行完毕==================")
	fmt.Println("==================任务调度器开始执行==================")
	tasks := []func(){
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("任务3完成")
		},
	}
	TaskScheduler(tasks)
	fmt.Println("==================任务调度器执行完毕==================")
}
