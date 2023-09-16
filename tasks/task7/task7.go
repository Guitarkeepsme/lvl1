// Реализовать конкурентную запись данных в map

package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu   sync.Mutex
	even map[int]bool
}

func newConturrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		even: make(map[int]bool),
	}
}

func (cm *ConcurrentMap) set(num int, parity bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.even[num] = parity
}

func (cm *ConcurrentMap) get(key int) (bool, bool) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	value, ok := cm.even[key]
	return value, ok
}

func main() {
	cm := newConturrentMap()

	cm.set(2, true)
	cm.set(3, false)
	value, ok := cm.get(3)
	if ok {
		fmt.Println(value) // Выведет: true
		fmt.Println(cm.even)
	}

}

// 	storage := make(map[int]bool)
// 	var mu sync.Mutex
// 	for i := 1; i <= 10; i++ {
// 		go func(i int) {
// 			if i%2 == 0 {
// 				storage[i] = true
// 			}
// 			if i%2 != 0 {
// 				mu.Lock()
// 				storage[i] = false
// 				mu.Unlock()
// 			}
// 		}(i)
// 	}
// 	fmt.Println(storage)
// }