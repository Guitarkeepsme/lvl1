/* Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false */

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Введите слово или фразу:")
	fmt.Scanln(&str)
	fmt.Printf("\nРезультат: %t\n\n", Unique(str))
}

func Unique(str string) bool {
	str = strings.ToLower(str)

	uniqueMap := make(map[string]int)

	for _, bar := range str {
		uniqueMap[string(bar)] += 1
		if uniqueMap[string(bar)] > 1 {
			return false
		}
	}
	return true
}
