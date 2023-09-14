// Описать все возможные способы остановки выполнения горутины.

// 1 способ: создание канала для передачи сигнала об остановки горутины
package main

import (
	"fmt"
	"time"
)

func worker(stop <-chan bool) {
	for {
		select {
		default:
			fmt.Println("Работаю...")
			time.Sleep(1 * time.Second)
		case <-stop:
			fmt.Println("Процесс остановлен!")
			return
		}
	}
}

func main() {
	stop := make(chan bool)

	go worker(stop)

	time.Sleep(5 * time.Second)
	stop <- true

	// Ждём завершения горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Программа завершена!")
}
