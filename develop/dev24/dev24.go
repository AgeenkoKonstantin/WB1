package main

import (
	"fmt"
	"math"
)

//24. Разработать программу нахождения расстояния между двумя точками, которые
//представлены в виде структуры Point с инкапсулированными параметрами x,y и
//конструктором

// Структура
type Point struct {
	x float32
	y float32
}

// Конструктор
func NewPoint(x, y float32) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Расстояние между двумя точками - корень из суммы квадратов разности координат
func Distance(p1, p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(
		math.Pow(float64(dx), 2) + math.Pow(float64(dy), 2),
	)
}
func main() {
	x := NewPoint(0.0, 5.0)
	y := NewPoint(5.0, 0.0)
	fmt.Println(Distance(x, y))
}
