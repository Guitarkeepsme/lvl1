// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Создаём структуру точки, чтобы
// присвоить координатам необходимые методы
type Point struct {
	x, y float64
}

// Этот метод
func (p *Point) SetPoint(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) Coords() (float64, float64) {
	return p.x, p.y
}

func CountAdress(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	var x1, y1, x2, y2 float64
	p1, p2 := new(Point), new(Point)

	fmt.Println("Задайте координаты первого числа:")
	_, err := fmt.Scanln(&x1, &y1)
	if err != nil {
		panic("Вы ввели некорректные данные!")
	}

	fmt.Println("Задайте координаты второго числа:")
	_, err = fmt.Scanln(&x2, &y2)
	if err != nil {
		panic("Вы ввели некорректные данные!")
	}

	p1.SetPoint(x1, y1)
	p2.SetPoint(x2, y2)

	res := CountAdress(*p1, *p2)

	fmt.Println("Расстояние между двумя точками:", res)
}
