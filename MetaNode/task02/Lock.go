package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
 * 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
 */
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

/**
 * 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
 * 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
 */
type AtomicCounter struct {
	count int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.count, 1)
}
func (c *AtomicCounter) GetCount() int {
	return int(atomic.LoadInt64(&c.count))
}

func main() {
	counter := SafeCounter{}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("最后计数: %d\n", counter.GetCount())
	counter1 := AtomicCounter{}
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter1.Increment()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("最后计数: %d\n", counter1.GetCount())
}
