package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	target := 12

	// Используем функцию sort.SearchInts() для бинарного поиска
	indexF := sort.SearchInts(arr, target)

	if indexF < len(arr) && arr[indexF] == target {
		fmt.Printf("Число %d найдено на позиции %d (с помощью sort.SearchInts())\n", target, indexF)
	} else {
		fmt.Printf("Число %d не найдено в массиве (с помощью sort.SearchInts())\n", target)
	}

	// Используем функцию sort.Search() для бинарного поиска
	indexS := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if indexS < len(arr) && arr[indexS] == target {
		fmt.Printf("Число %d найдено на позиции %d (с помощью sort.Search())\n", target, indexS)
	} else {
		fmt.Printf("Число %d не найдено в массиве (с помощью sort.Search())\n", target)
	}
}
