// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func binarySearch(arr []int, tgt int) (i int, ok bool) {
	// Определяем первый и последний элементы
	left := 0
	right := len(arr) - 1
	res := 0
	// Пока первый ниже или равен последнему
	for left <= right {
		// Серединой будет это:
		mdl := (left + right) / 2
		// И если мы уже нашли таргет,
		// заносим его в результат
		if arr[mdl] == tgt {
			res = mdl
		}
		// Если мы оказались слева, сдвигаем первый индекс
		// вправо
		if arr[mdl] < tgt {
			left = mdl + 1
			// Иначе наоборот
		} else {
			right = mdl - 1
		}
	}
	// Если мы дошли до конца, но так и не нашли таргет,
	// это значит, таргета тут нет
	if left == len(arr) || arr[right] != tgt {
		return 0, false
	}
	return res, true
}

func main() {
	fmt.Println(binarySearch([]int{3, 5, 7, 10, 13, 20}, 75))
}
