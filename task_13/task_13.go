package main

import (
	"fmt"
)

// Метод 1: Простая реализация через a, b = b, a
func method1() {
	a, b := 5, 10
	fmt.Println("Метод 1:")
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}

// Метод 2: Сложение и вычитание
func method2() {
	a, b := 5, 10
	fmt.Println("Метод 2:")
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}

// Метод 3: Исключающее ИЛИ (XOR)
func method3() {
	a, b := 5, 10
	fmt.Println("Метод 3:")
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}

// Метод 4: Использование указателей
func method4() {
	a, b := 5, 10
	fmt.Println("Метод 4:")
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	swapByPointers(&a, &b)

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}

// Функция для обмена значениями через указатели
func swapByPointers(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	method1()
	fmt.Println()
	method2()
	fmt.Println()
	method3()
	fmt.Println()
	method4()
	fmt.Println()
}
