/* Написать программу, которая конкурентно рассчитает значение квадратов
чисел, взятых из массива (2,4,6,8,10), и выведет их квадраты в stdout. */

package main

import (
	"fmt"
	"os"
	"sort"
)

// Определение функции возведения в квадрат
func square(num int, c chan int) {
	c <- num * num // Запись результата в канал
}

// Определение главной функции
func main() {
	// Создание канала
	c := make(chan int)
	// Перменная из списка (среза?) чисел
	nums := []int{2, 4, 6, 8, 10}
	for _, num := range nums {
		go square(num, c) // Запуск горутин
	}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		result = append(result, <-c) // Добавление в массив при получении результата из канала
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j] // Сортировка массива
	})

	fmt.Fprintln(os.Stdout, result) // Вывод результата в stdout
}
