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
	workersCount := flag.Int("workersCount", 1, "Выберите количество воркеров")
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

	// Главный поток записывает данные в канал
	wg.Add(1)
	go func(ctx context.Context, dataChan chan int) {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done(): // Если контекст завершён,
				close(dataChan) // закрываем канал
				return
			default:
				dataChan <- i
			}
		}
	}(ctx, dataChan)

	closeChan := make(chan os.Signal, 1)
	// Ждём сигнала Ctrl+C о завершении работы и передаём его в канал закрытия
	signal.Notify(closeChan, syscall.SIGINT)
	<-closeChan                                          // Пока не нажали Ctrl+C, канал блокируется на этой строчке
	fmt.Printf("Команда получена, завершаю работу...\n") // Как только команда пришла, печатается уведомление о ней
	cancel()                                             // И завершается контекст
	wg.Wait()                                            // Затем начинается ожидания закрытия воркеров

	fmt.Printf("Воркеры завершили работу, отключаюсь.\n") // Печатаем результат завершения программы
}

// Создание воркера: он принимает контекст, айди для отображения в консоли, числовые данные из канала,
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
