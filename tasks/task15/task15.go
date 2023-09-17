/* К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	someFunc()
	rightFunc()
}

// Исходя из этого фрагмента кода, createHugeString
// умножает исходную строку на заданное число раз.
// Напишем эту функцию:

func createHugeString(mult int) string {
	var str strings.Builder

	defStr := "word "
	word := flag.String("word", defStr, "write something")
	flag.Parse()

	for i := 0; i < mult; i++ {
		str.WriteString(*word)
	}
	return str.String()
}

// В этой части кода мы применяем предыдущую функцию,
// а затем из полученной огромной строки забираем
// последние 100 символов
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]

	// Распечатаем результат для наглядности
	fmt.Println(justString)
}

// Если мы запустим программу с дефолтным значением слова, она будет работать корректно.
// Однако как только мы заменим это слово на набор символов, например, кириллицы
// (или другого юникода), мы увидим, что часть слов стёрлась. Дело в том, что юникод-символы
// занимают 2 байта вместо одного, поэтому при такой реализации функции часть информации пропадёт.

// Нужно конвертировать эти символы в руны:

func rightFunc() {
	v := createHugeString(1 << 10)
	rn := []rune(v)
	justString = string(rn[:100])

	// Распечатаем результат для наглядности
	fmt.Println(justString)
}
