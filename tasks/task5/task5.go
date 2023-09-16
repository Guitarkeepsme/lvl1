/*
	Разработать программу, которая будет последовательно отправлять значения в канал,

а с другой стороны канала — читать. По истечении N секунд программа должна завершаться.
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Устанавливаем флаг на количество секунд
	seconds := flag.Int("seconds", 5, "set amount of time")
	flag.Parse()

	// Вызываем контекст, который сменится после вышеозначенного времени
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*seconds)*time.Second)
	defer cancel()

	// Создаём канал, который...
	infoChan := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(ctx context.Context, infoChan chan int) {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				close(infoChan)
				fmt.Println("Время вышло!")
				return
			default:
				infoChan <- i
			}
		}
	}(ctx, infoChan)

	go reader(wg, infoChan)
	wg.Wait()

}

func reader(wg *sync.WaitGroup, infoChan <-chan int) {
	defer wg.Done()
	// Ридер работает в бесконечном цикле:
	for data := range infoChan {
		fmt.Printf("Читаю данные %d\n", data) // И это отображается в консоли
	}
}
