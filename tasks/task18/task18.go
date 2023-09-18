// Реализовать структуру-счетчик,
// которая будет инкрементироваться в конкурентной среде.
// По завершении программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	mutex sync.Mutex
	num   int64
}

func numCounter() *Counter {
	return &Counter{
		mutex: sync.Mutex{},
	}
}

func (c *Counter) MutexCounter() {
	c.mutex.Lock()
	c.num++
	c.mutex.Unlock()
}

func (c *Counter) AtomicCounter() {
	atomic.AddInt64(&c.num, 1)
}

func main() {
	defNum := 10
	cnt1, cnt2 := numCounter(), numCounter()

	for i := 0; i < defNum; i++ {
		go func() {
			for {
				cnt1.MutexCounter()
			}
		}()
	}

	for i := 0; i < defNum; i++ {
		go func() {
			for {
				cnt2.AtomicCounter()
			}
		}()
	}

	time.Sleep(10 * time.Second)

	fmt.Println("Посчитано через Mutex:  ", cnt1.num)
	fmt.Println("Посчитано через Atomic: ", cnt2.num)
}
