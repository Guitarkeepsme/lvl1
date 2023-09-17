// Реализовать структуру-счетчик,
// которая будет инкрементироваться в конкурентной среде.
// По завершении программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	defNum := 5000
	cnt1, cnt2 := numCounter(), numCounter()

	for i := 0; i < defNum; i++ {
		go cnt1.MutexCounter()
	}

	for i := 0; i < defNum; i++ {
		go cnt2.AtomicCounter()
	}

	fmt.Println("Посчитано через Mutex:  ", cnt1.num)
	fmt.Println("Посчитано через Atomic: ", cnt2.num)
}
