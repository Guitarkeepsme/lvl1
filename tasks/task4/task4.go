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
	workerCount := flag.Int("workerCount", 1, "set count of workers")
	flag.Parse()

	// Добавляем контекст: если возникает отмена, программа останавливается
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаём канал передачи данных
	dataChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(*workerCount)

	for i := 0; i < *workerCount; i++ {
		go worker(ctx, i, dataChan, &wg)
	}

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
	signal.Notify(closeChan, syscall.SIGINT)
	<-closeChan
	fmt.Printf("Got SIGTERM, exiting...\n")
	cancel()
	close(dataChan)
	wg.Wait()
	fmt.Printf("All workers done, exiting...\n")
}

func worker(ctx context.Context, id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worder %d\n", id)
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-dataChan:
			fmt.Printf("Worker %d: %d\n", id, data)
		}
	}
}
