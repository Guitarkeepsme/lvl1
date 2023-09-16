/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode. */

package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Выставляем флаг на выбор слова, которое нужно перевернуть
	defStr := "главрыба"
	str := flag.String("word", defStr, "Введите слово")
	flag.Parse()

	// Переводим строку в слайс рун
	runes := []rune(*str)

	// Результатом будет строка, которую мы соберём
	// с помощью strings.Builder
	res := strings.Builder{}

	// Проходим по каждому символу слайса рун
	// от последнего к первому
	for i := len(runes) - 1; i >= 0; i-- {
		res.WriteRune(runes[i])
	}

	fmt.Println(res.String())
}
