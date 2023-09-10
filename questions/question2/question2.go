// Интерфейс -- это конструкция, которая нужна для того, чтобы определять и описывать некие методы, необходимые какому-то типу данных.

package main

import (
	"fmt"
	"reflect"
)

// Пусть интерфейс Phone имеет сигнатурное значение getInfo:
type Phone interface {
	GetInfo() string
}

// А также существуют типы Iphone12 и SamsungGalaxyS20:
type Iphone12 struct{}
type GalaxyS20 struct{}

// Эти типы данных удовлетворяют интерфейсу Phone, поскольку имеют сигнатурное значение getInfo:
func (Iphone12) GetInfo() string {
	return "Iphone 12 is made by Apple"
}
func (GalaxyS20) GetInfo() string {
	return "Galaxy S20 is made by Samsung"
}

// Благодаря этому, если мы определим функцию printInfo, которая принимает тип данных Phone,
func printInfo(p Phone) {
	fmt.Println(p.GetInfo())
}

// Мы сможем использовать эту функцию для типов Iphone12 и GalaxyS20,
// Ведь они также относятся к типу данных Phone благодаря удовлетворяемости одноимённому интерфейсу.
func main() {
	ip := Iphone12{}
	sg := GalaxyS20{}

	printInfo(ip) // Вывод: Iphone 12 is made by Apple
	printInfo(sg) // Вывод: Galaxy S20 is made by Samsung

	a := "a"
	b := "1"
	emptyInterface(a)
	emptyInterface(b)
}

// Помимо этого, интерфейс может быть пустым. В таком случае он принимает любой тип данных:

func emptyInterface(i interface{}) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.String:
		fmt.Println(i.(string))
	case reflect.Int:
		fmt.Printf("%d", i.(int))
	}
}
