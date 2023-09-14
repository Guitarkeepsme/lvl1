/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Инициализируем флаг, позволяющий выбрать количество воркеров
	workersCount := flag.Int("workersCount", 1, "set count of workers")
	flag.Parse()

	// Добавляем контекст: если возникает отмена, программа останавливается
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаём канал передачи данных
	dataChan := make(chan int)

	// Создаём WaitGroup для ожидания завершения работы всех воркеров
	var wg sync.WaitGroup
	wg.Add(*workersCount)
	// Запускаем горутину для каждого воркера
	for i := 0; i < *workersCount; i++ {
		go worker(ctx, i, dataChan, &wg)
	}

	// ПОЧЕМУ СЮДА НУЖЕН КОНТЕКСТ? Не совсем понимаю принцип работы этой части кода

	// Главный поток записывает данные в канал
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				dataChan <- i
			}
		}
	}(ctx)

	closeChan := make(chan os.Signal, 1)
	// Ждём сигнала Ctrl+C о завершении работы и передаём его в канал закрытия
	signal.Notify(closeChan, syscall.SIGINT)
	<-closeChan
	fmt.Printf("Got SIGINT, exiting...\n") // Печатаем информацию о подачи сигнала завершения
	cancel()                               // Когда вызывается эта функция, воркер отрабатывает первый кейс и завершает работу
	close(dataChan)                        // Закрываем канал

	wg.Wait()                                    // Ждём синхронизации закрытия воркеров
	fmt.Printf("All workers done, exiting...\n") // Печатаем результат завершения программы
}

// Создание воркера: он принимает контекст, айди, числовые данные из канала,
// а также информацию из WaitGroup
func worker(ctx context.Context, id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()               // Откладываем завершение его работы
	fmt.Printf("worker %d\n", id) // Печатаем процесс работы
	// Воркер работает в бесконечном цикле:
	for {
		select {
		case <-ctx.Done(): // Либо он завершает работу
			return
		case data := <-dataChan: // Либо читает данные из канала
			fmt.Printf("Worker %d: %d\n", id, data) // И это отображается в консоли
		}
	}
}
