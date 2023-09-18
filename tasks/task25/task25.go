// Реализовать собственную функцию sleep.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ждём 5 секунд...")
	SleepByChan(5 * time.Second)
	fmt.Println("Готово!")

	fmt.Println("Ждём 4 секунды...")
	SleepByTicker(4 * time.Second)
	fmt.Println("Готово!")

	fmt.Println("Ждём 3 секунды...")
	SleepByContext(3 * time.Second)
	fmt.Println("Готово!")
}

func SleepByChan(d time.Duration) {
	<-time.After(d)
}

func SleepByTicker(d time.Duration) time.Time {
	ticker := time.Tick(d)
	for wake := range ticker {
		return wake
	}
	return time.Now()
}

func SleepByContext(d time.Duration) {
	ctx, _ := context.WithTimeout(context.Background(), d) // Здесь отмена не нужна
	<-ctx.Done()
}
