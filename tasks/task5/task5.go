/*
	Разработать программу, которая будет последовательно отправлять значения в канал,

а с другой стороны канала — читать. По истечении N секунд программа должна завершаться.
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

func main() {
	seconds := flag.Int("seconds", 5, "set amount of time")
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*seconds)*time.Second)
	defer cancel()

	infoChan := make(chan int)

	go func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				infoChan <- i
			}
		}
	}(ctx)

	for i := 0; ; i++ {
		go reader(ctx, infoChan)
	}
}

func reader(ctx context.Context, infoChan <-chan int) {
	// Ридер работает в бесконечном цикле:
	for {
		select {
		case <-ctx.Done(): // Либо он завершает работу
			return
		case data := <-infoChan: // Либо читает данные из канала
			fmt.Printf("Читаю данные %d:\n", data) // И это отображается в консоли
		}
	}
}
