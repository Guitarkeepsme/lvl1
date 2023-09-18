/* Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false otherwise). Функция проверки должна быть регистронезависимой.

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
	// Определяем переменную, в которую
	// впишем строку через консоль
	var str string
	fmt.Println("Введите слово или фразу:")
	fmt.Scanln(&str)
	fmt.Printf("\nРезультат: %t\n\n", isUnique(str))
}

func isUnique(str string) bool {
	// Прежде всего переводим символы в нижний регистр,
	// чтобы соблюсти условие задачи
	str = strings.ToLower(str)

	// Создаём карту, в которую будем вносить данные
	// по каждому элементу строки
	uniqueMap := make(map[string]int)

	// Циклом проходим по всем символам строки
	for _, bar := range str {
		// и считаем их: если будет несколько одинаковых,
		// их значение станет выше единицы,
		uniqueMap[string(bar)] += 1
		// и тогда мы вернём false
		if uniqueMap[string(bar)] > 1 {
			return false
		}
	}
	// Иначе возвращаем true
	return true
}
