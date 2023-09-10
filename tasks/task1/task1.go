/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import (
	"fmt"
	"math/rand"
)

// Объявление структуры Human
type Human struct {
	// Примерный набор полей: имя, фамилия, возраст
	firstName string
	lastName  string
	Age       int
}

// Пример метода
func (h *Human) SaySomething() string {
	return "Wish me luck!\n"
}
func (h *Human) DiceThrow() int {
	min := 2
	max := 12
	return rand.Intn((max - min) + min)
}

// Объявление структуры Action
type Action struct {
	// Встраивание методов Human в структуру Action
	Human
}

func main() {
	// Сначала прописываем значения полей Action
	Andrew := Action{
		Human: Human{
			firstName: "Andrew",
			lastName:  "Black",
			Age:       25,
		},
	}
	// Результат
	fmt.Printf("%+s", Andrew.SaySomething())
	fmt.Println(Andrew.DiceThrow())

}
