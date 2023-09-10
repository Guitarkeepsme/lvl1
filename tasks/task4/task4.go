/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Сначала создаём канал, который будет передавать данные между воркерами
	dataChan := make(chan int)

	// Указываем количество воркеров, например 3
	workerCount := 3

	// Создаём WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup
	wg.Add(workerCount)

	// Запускаем воркеров и проходим по ним циклом
	for i := 0; i < workerCount; i++ {
		go worker(i, dataChan, &wg)
	}

	// Главный поток записывает данные в канал
	go func() {
		for i := 0; ; i++ {
			dataChan <- i
		}
	}()

	// Наконец, ждём, когда все воркеры завершат работу
	wg.Wait()
}

// Функция, создающая воркеров
func worker(id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Читаем данные из канала и выводим их в stdout
	for data := range dataChan {
		fmt.Printf("Worker %d: %d\n", id, data)
	}
}
