/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

// Объявление структуры Human
type Human struct {
	// Примерный набор полей: имя, фамилия, возраст
	firstName string
	lastName  string
	Age       int
}

// Объявление структуры Action
type Action struct {
	// Примерный набор полей: что-то говорит (если говорит, то что именно), продолжает молчать (да или нет)
	saysSomething string
	remainsSilent bool
	// Встраивание методов Human в структуру Action
	Human
}

func main() {
	// Сначала прописываем значения полей Action
	Andrew := Action{
		saysSomething: "Hi everyone!",
		remainsSilent: false,
		// Затем прописываем значения полей Human
		Human: Human{
			firstName: "Andrew",
			lastName:  "Black",
			Age:       25,
		},
	}

	// Распечатывание результата
	fmt.Printf("%+v", Andrew)

}
