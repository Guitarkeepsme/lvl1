// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

// func gen(nums []int) <- chan int{

// }

func main() {
	// НУЖЕН ЛИ ТУТ МЬЮТЕКС?
	var mutex sync.Mutex

	nums := []int{3, 7, 20, 31}

	// Создаём два буферизированных канала: в один будем писать числа из nums,
	// В другой фиксируем результат возведения этих чисел в квадрат.

	// БЕЗ БУФЕРА происходил дедлок. Почему?
	startChan := make(chan int, len(nums))
	resChan := make(chan int, len(nums))

	// Проходим циклом по массиву и отправляем каждое число в канал
	for _, num := range nums {
		mutex.Lock()
		startChan <- num
		mutex.Unlock()
	}

	// Закрываем канал после того, как перенесли в него всю информацию
	close(startChan)

	// Теперь циклом возводим каждое число из первого канала в квадрат
	// и отправляем полученный результат во второй канал
	for num := range startChan {
		mutex.Lock()
		resChan <- num * num
		mutex.Unlock()
	}

	// Закрываем второй канал
	close(resChan)

	// Печатаем результат
	for square := range resChan {
		fmt.Println(square)
	}

}
