// Интерфейс -- это конструкция, которая нужна для того, чтобы определять и описывать некие методы, необходимые какому-то типу данных.

package main

import "fmt"

// Пусть интерфейс Phone имеет сигнатурное значение getInfo:
type Phone interface {
	getInfo() string
}

// А также существуют типы Iphone12 и SamsungGalaxyS20:
type Iphone12 struct{}
type GalaxyS20 struct{}

// Эти типы данных удовлетворяют интерфейсу Phone, поскольку имеют сигнатурное значение getInfo:
func (Iphone12) getInfo() string {
	return "Iphone 12 is made by Apple"
}
func (GalaxyS20) getInfo() string {
	return "Galaxy S20 is made by Samsung"
}

// Благодаря этому, если мы определим функцию printInfo, которая принимает тип данных Phone,
func printInfo(p Phone) {
	fmt.Println(p.getInfo())
}

// Мы сможем использовать эту функцию для типов Iphone12 и GalaxyS20,
// Ведь они также относятся к типу данных Phone благодаря удовлетворяемости одноимённому интерфейсу.
func main() {
	ip := Iphone12{}
	sg := GalaxyS20{}

	printInfo(ip) // Вывод: Iphone 12 is made by Apple
	printInfo(sg) // Вывод: Galaxy S20 is made by Samsung
}
