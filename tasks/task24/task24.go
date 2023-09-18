// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

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
	p1, p2 := new(Point), new(Point)

	p1.SetPoint(2, 3)

	p2.SetPoint(4, 3)

	res := CountAdress(*p1, *p2)

	fmt.Println("Расстояние между двумя точками:", res)
}
