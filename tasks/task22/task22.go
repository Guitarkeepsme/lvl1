// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"flag"
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	defA := "283904503245"
	defB := "9896954876927"

	a := flag.String("a", defA, "Введите число а")
	b := flag.String("b", defB, "Введите число b")
	flag.Parse()

	parsedA, _ := decimal.NewFromString(*a)
	parsedB, _ := decimal.NewFromString(*b)

	sum := parsedA.Add(parsedB)
	sub := parsedA.Sub(parsedB)
	mul := parsedA.Mul(parsedB)
	div := parsedA.Div(parsedB)

	oper := flag.String("operation", sum.String(), "Выберите операцию: \n1 - Сложениe\t2 - Вычитание\n3 - Умножение\t4 - Деление\n")
	sumF := flag.String("sum", sum.String(), "Сумма двух чисел")
	subF := flag.String("sub", sub.String(), "Разность двух чисел")
	mulF := flag.String("mul", mul.String(), "Умножение двух чисел")
	divF := flag.String("div", div.String(), "Деление двух чисел")
	flag.Parse()

	fmt.Printf("%s\n", *sumF)
	fmt.Printf("%s\n", *subF)
	fmt.Printf("%s\n", *mulF)
	fmt.Printf("%s\n", *divF)
}
