// Реализовать собственную функцию sleep.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Демонстрация работы функции в трёх воплощениях
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

// 1 способ: с помощью After
func SleepByChan(d time.Duration) {
	// After выжидает, пока истечёт d,
	// а затем отправляет данные в главный канал
	<-time.After(d)
}

// 2 способ: через тикер
func SleepByTicker(d time.Duration) time.Time {
	// Тикер засекает d времени
	ticker := time.Tick(d)

	// И здесь по прохождим по каждой единице времени
	// и возвращаем wake
	for wake := range ticker {
		return wake
	}
	// После того, как тикер завершился,
	// возвращаем время
	return time.Now()
}

// 3 способ: с помощью контекста
func SleepByContext(d time.Duration) {
	// Ставим контекст с таймаутом: через d он сменится
	ctx, _ := context.WithTimeout(context.Background(), d) // Здесь отмена не нужна
	// возвращаем смену контекста в главный канал
	<-ctx.Done()
}
