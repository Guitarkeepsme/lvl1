// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import "fmt"

func partition(arr []int, first, last int) ([]int, int) {
	pivot := arr[last]

	i := first
	for n := first; n < last; n++ {
		if arr[i] < pivot {
			arr[i], arr[n] = arr[n], arr[i]
			i++
		}
	}
	arr[i], arr[last] = arr[last], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	fmt.Println(quickSortStart([]int{1, 3, 5, 2, 10, 7, 35, 0}))
}
