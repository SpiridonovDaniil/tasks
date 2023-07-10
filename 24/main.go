package main

import (
	"fmt"
	"math"
)

type WorkWithPoint interface { //интерфейс описывающий как работать со структурой Point.
	ReadCoordinates() (x, y float64)
}

type Point struct { //структура с инкапсулированными полями.
	x float64
	y float64
}

func NewPoint(x, y float64) *Point { //конструктор структуры Point.
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) ReadCoordinates() (x, y float64) { //метод для получения данных из структуры Point.
	return p.x, p.y
}

func main() {
	firstPoint := NewPoint(6, 7) //инициализируем две точки.
	secondPoint := NewPoint(10, 10)
	fmt.Println(distanceCalculate(firstPoint, secondPoint))
}

func distanceCalculate(point1, point2 *Point) float64 { //функция подсчета расстояния между двумя точками.
	x1, y1 := point1.ReadCoordinates() //считываем данные из структуры Point.
	x2, y2 := point2.ReadCoordinates()

	x := x2 - x1
	y := y2 - y1

	d := math.Hypot(x, y)

	return d
}
