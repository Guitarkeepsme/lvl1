// Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func setBit(n int64, pos uint, value int) int64 {
	mask := int64(1) << pos
	if value == 1 {
		n |= mask
	} else {
		n |= ^mask
	}
	return n
}

func main() {
	var num int64
	var i, value int

	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	fmt.Print("Введите индекс бита: ")
	fmt.Scanln(&i)

	fmt.Print("Введите значение: 1 или 0: ")
	fmt.Scanln(&value)

	res := setBit(num, uint(i), value)
	fmt.Printf("Результат: %d\n", res)
}
