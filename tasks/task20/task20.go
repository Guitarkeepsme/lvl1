/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow» */

package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Выставляем флаг на выбор фразы, которую нужно изменить
	defStr := "snow dog sun"
	str := flag.String("sentence", defStr, "Введите слово")
	flag.Parse()

	// Переводим строку в слайс рун
	runes := string([]rune(*str))

	// Убираем пробелы
	splited := strings.Split(string(runes), " ")

	// Переменная ss нужна для сбора слов в обратном порядке
	var ss []string

	// Проходимся по списку слов в обратном порядке
	for i := len(splited) - 1; i >= 0; i-- {
		ss = append(ss, splited[i])
	}
	// Соединяем слова в одну строку
	res := strings.Join(ss, " ")

	fmt.Println(res)

}
