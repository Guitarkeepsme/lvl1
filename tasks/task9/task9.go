// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

// Создаём две функции, которые будут запускаться в горутинах:
// первая принимает числа из канала, возводит их в квадрат и
// отправляет во второй канал:
func pipelineMultiplier(in chan int, out chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

// Вторая принимает информацию о количестве ожидающих горутин
// и квадраты чисел, а затем выводит их в консоль
func pipelinePrinter(wg *sync.WaitGroup, out chan int) {
	for el := range out {
		fmt.Println(el)
	}
	// После этого все ожидающие горутины завершаются
	// и размораживается main
	wg.Done()
}

func main() {

	nums := []int{3, 7, 20, 31, 9, 1356}

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Создаём два буферизированных канала: в один будем писать числа из nums,
	// В другой фиксируем результат возведения этих чисел в квадрат.
	startChan := make(chan int, len(nums))
	resChan := make(chan int, len(nums))

	go func() {
		for _, num := range nums {
			startChan <- num
		}

		close(startChan)
		wg.Done()
	}()
	// Запускаем
	go pipelineMultiplier(startChan, resChan)
	go pipelinePrinter(&wg, resChan)

	wg.Wait()
}
