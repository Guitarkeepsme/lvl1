// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func binarySearch(arr []int, tgt int) (i int, ok bool) {
	first := 0
	last := len(arr) - 1
	res := 0
	for first <= last {
		mdl := (first + last) / 2

		if arr[mdl] == tgt {
			res = mdl
		}
		if arr[mdl] < tgt {
			first = mdl + 1
		} else {
			last = mdl - 1
		}
	}
	if first == len(arr) || arr[first] != tgt {
		return 0, false
	}
	return res, true
}

func main() {
	fmt.Println(binarySearch([]int{3, 5, 7, 10, 13, 20}, 75))
}
