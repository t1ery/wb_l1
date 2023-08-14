package main

import (
	"fmt"
)

// Метод 1: Использование map с пустыми значениями
func method1() {
	fmt.Println("Метод 1:")
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueSet := make(map[string]struct{})

	for _, item := range sequence {
		uniqueSet[item] = struct{}{}
	}

	fmt.Println("Уникальные элементы:")
	for item := range uniqueSet {
		fmt.Println(item)
	}
}

// Метод 2: Использование map с булевыми значениями
func method2() {
	fmt.Println("Метод 2:")
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueSet := make(map[string]bool)

	for _, item := range sequence {
		uniqueSet[item] = true
	}

	fmt.Println("Уникальные элементы:")
	for item := range uniqueSet {
		fmt.Println(item)
	}
}

// Метод 3: Использование среза для хранения уникальных значений
func method3() {
	fmt.Println("Метод 3:")
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueSet := []string{}

	for _, item := range sequence {
		found := false
		for _, uniqueItem := range uniqueSet {
			if item == uniqueItem {
				found = true
				break
			}
		}
		if !found {
			uniqueSet = append(uniqueSet, item)
		}
	}

	fmt.Println("Уникальные элементы:")
	for _, item := range uniqueSet {
		fmt.Println(item)
	}
}

// Метод 4: Использование структуры Set (пользовательский тип, по сути реализации как 2 метод)
type Set map[string]bool

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(item string) {
	s[item] = true
}

func method4() {
	fmt.Println("Метод 4:")
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueSet := NewSet()

	for _, item := range sequence {
		uniqueSet.Add(item)
	}

	fmt.Println("Уникальные элементы:")
	for item := range uniqueSet {
		fmt.Println(item)
	}
}

func main() {
	method1()
	fmt.Println()
	method2()
	fmt.Println()
	method3()
	fmt.Println()
	method4()
}
