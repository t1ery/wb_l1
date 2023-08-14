package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Способы, которые различаются:

// 1. Композиция с интерфейсами
type ShapeAdapter1 struct {
	Shape
}

// 2. Использование функций высшего порядка
type ShapeAdapter3 struct {
	AreaFunc func() float64
}

// 3. Использование интерфейсов с пустыми реализациями
type ShapeAdapter4 struct {
	Rectangle
}

func (s ShapeAdapter4) Area() float64 {
	return s.Rectangle.Area()
}

// Способы, которые похожи (используют встраивание):

// 4. Встраивание в анонимное поле
type ShapeAdapter2 struct {
	Rectangle
}

// 5. Использование наследования
type ShapeAdapter5 struct {
	Rectangle
}

func main() {
	rectangle := Rectangle{Width: 5, Height: 10}

	// 1. Композиция с интерфейсами
	adapter1 := ShapeAdapter1{Shape: rectangle}
	fmt.Printf("1. Площадь через композицию: %.2f\n", adapter1.Area())

	// 3. Использование функций высшего порядка
	adapter3 := ShapeAdapter3{AreaFunc: rectangle.Area}
	fmt.Printf("3. Площадь через функцию высшего порядка: %.2f\n", adapter3.AreaFunc())

	// 4. Использование интерфейсов с пустыми реализациями
	adapter4 := ShapeAdapter4{Rectangle: rectangle}
	fmt.Printf("4. Площадь через интерфейсы с пустыми реализациями: %.2f\n", adapter4.Area())

	// 2. Встраивание в анонимное поле
	adapter2 := ShapeAdapter2{Rectangle: rectangle}
	fmt.Printf("2. Площадь через встраивание в анонимное поле: %.2f\n", adapter2.Area())

	// 5. Использование наследования
	adapter5 := ShapeAdapter5{Rectangle: rectangle}
	fmt.Printf("5. Площадь через наследование: %.2f\n", adapter5.Area())
}
