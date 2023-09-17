// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}

	i = "ldkfgjd"

	// fmt.Print("Введите значение, тип которого надо проверить: \n\n")
	// fmt.Scanln(&i)

	res := getType(i)
	fmt.Printf("Тип переменной: %s\n", res)
}

func getType(i interface{}) string {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Int:
		return "Int"
	case reflect.String:
		return "String"
	case reflect.Bool:
		return "Boolean"
	case reflect.Chan:
		return "Channel"
	default:
		return "Неизвестный тип"
	}
}
