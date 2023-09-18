// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"flag"
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// В качестве дефолтных значений берём величины, явно превышающие 2^20
	defA := "28390450324533"
	defB := "9896954876927234"

	// Инициализируем флаги для выбора кастомных значений
	a := flag.String("a", defA, "Введите число а")
	b := flag.String("b", defB, "Введите число b")
	// И флаг для операции
	oper := flag.String("oper", "sum", "Выберите операцию: \nsum - Сложениe\tsub - Вычитание\nmul - Умножение\tdiv - Деление\n")
	flag.Parse()
	// Теперь переводим величины в тип данных decimal для дальнейших вычислений
	parsedA, _ := decimal.NewFromString(*a)
	parsedB, _ := decimal.NewFromString(*b)

	// Проводим операции с ними

	// Создаём условия для выбора операции
	switch *oper {
	case "sum":
		sum := parsedA.Add(parsedB)
		fmt.Printf("%s\n", sum)
	case "sub":
		sub := parsedA.Sub(parsedB)
		fmt.Printf("%s\n", sub)
	case "mul":
		mul := parsedA.Mul(parsedB)
		fmt.Printf("%s\n", mul)
	case "div":
		div := parsedA.Div(parsedB)
		fmt.Printf("%s\n", div)
	default:
		panic("Введена неизвестная команда!")
	}

}
