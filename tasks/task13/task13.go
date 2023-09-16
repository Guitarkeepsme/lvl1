// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите два числа: ")
	fmt.Scan(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("Теперь a = %d, b = %d\n", a, b)
}
