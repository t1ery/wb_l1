package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64 // Начинаются с маленькой буквы - инкапсулированные поля
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	// Вычисляем Евклидово расстояние
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	distance := point1.Distance(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
