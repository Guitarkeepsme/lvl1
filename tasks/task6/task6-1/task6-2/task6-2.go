// Описать все возможные способы остановки выполнения горутины.

// 2 способ: использовать контекст

package main

import (
	"context"
	"time"
)

func main() {
	stop := make(chan bool)

	go worker(stop)

	time.Sleep(5 * time.Second)
	stop <- true
}

func worker(ctx context.Context) {

}
