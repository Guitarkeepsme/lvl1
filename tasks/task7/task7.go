// Реализовать конкурентную запись данных в map

package main

import (
	"fmt"
	"sync"
)

// Создаём структуру с картой
// и мьютексом, нужным для конкурентной работы карты
type ConcurrentMap struct {
	mu   sync.RWMutex
	even map[int]int
}

func newConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		even: make(map[int]int),
	}
}

func (cm *ConcurrentMap) set(num int, parity int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.even[num] = parity
}

func (cm *ConcurrentMap) get(key int) (int, bool) {
	cm.mu.RLock()
	defer cm.mu.Unlock()
	value, ok := cm.even[key]
	return value, ok
}

func main() {
	cm := newConcurrentMap()

	wg := sync.WaitGroup{}

	go func() {
		for i := 1; i <= 10000; i++ {
			if i%2 == 0 {
				cm.set(i, i)
			} else {
				cm.set(i, i)
			}
		}
	}()

	go func() {
		for i := 1; i <= 10000; i++ {
			if i%2 == 0 {
				cm.set(i, i)
			} else {
				cm.set(i, i+100000)
			}
		}
	}()

	wg.Wait()

	for _, char := range cm.even {
		fmt.Println(char)
	}
	value, ok := cm.get(3)
	if ok {
		fmt.Println(value) // Выведет: true
		fmt.Println(cm.even)
	}

}
