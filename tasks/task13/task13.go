// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		panic("Вы ввели некорректные данные!")
	}
	fmt.Printf("a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("Теперь a = %d, b = %d\n", a, b)
}
