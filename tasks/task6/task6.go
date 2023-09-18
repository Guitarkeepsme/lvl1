// Описать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func stopByChan() {
	// Инициализируем канал для остановки горутины
	stop := make(chan bool)

	// Вызываем анонимную горутину
	go func() {
		for {
			select {
			// Когда в канал перейдёт true, останавливаем процесс
			case <-stop:
				fmt.Println("Канал закрыт.")
				return
			// Иначе демонстрируем работу
			default:
				fmt.Println("Работаю...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	// Ставим счётчик для наглядности работы
	time.Sleep(3 * time.Second)
	// По завершении отчёта отправляем true для закрытия канала
	stop <- true

	// Ждём завершения горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Программа завершена после закрытия канала!")
}

// 2 способ: остановка горутины с использованием контекста

func stopByContext(ctx context.Context) {
	// Выставляем контекст с таймаутом в 4 секунды
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	// Откладываем отмену контекста
	defer cancel()
	// Инициализируем канал для передачи информации о контексте
	check := make(chan bool)
	// Создаём анонимную горутину
	go func() {
		for {
			select {
			// Если контекст завершён, отправляем в канал true и сообщаем об этом в консоли
			case <-ctx.Done():
				check <- true
				fmt.Println("Программа завершена после смены контекста!")
				return
			// Иначе демонстрируем работу
			default:
				fmt.Println("Работаю...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	// Читаем данные канала
	<-check

}

// 3 способ: остановка горутины после паники и обработка паники с помощью recover

func stopByPanic() {
	defer func() {
		// Обрабатываем панику
		if err := recover(); err != nil {
			log.Println("Произошла паника:", err)
		}
		time.Sleep(1 * time.Second)
		// Демонстрируем жизнь после паники
		fmt.Println("Живу после паники!")
	}()
	// Вызываем панику делением на ноль
	a, b := 1, 0
	fmt.Println(a / b)
}

func main() {
	go stopByPanic() // Запускаем горутину
	time.Sleep(2 * time.Second)
	stopByContext(context.Background())
	time.Sleep(2 * time.Second)
	stopByChan()
}
