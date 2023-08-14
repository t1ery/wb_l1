package main

import (
	"fmt"
	"sort"
)

func main() {
	// Исходный массив
	arr := []int{12, 4, 5, 6, 7, 3, 1, 15, 9, 8}
	fmt.Println("Исходный массив:", arr)

	// Используем встроенный метод сортировки для быстрой сортировки
	sort.Ints(arr)

	fmt.Println("Отсортированный массив:", arr)
}
