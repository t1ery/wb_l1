package main

import (
	"fmt"
	"sort"
)

// Метод с использованием карт (Maps)
func intersectionWithMap(slice1, slice2 []int) []int {

	// Преобразуем срезы в мапы для быстрого доступа
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, elem := range slice1 {
		set1[elem] = true
	}
	for _, elem := range slice2 {
		set2[elem] = true
	}

	var result []int

	// Перебираем элементы первой мапы
	for elem := range set1 {
		// Если элемент также есть во второй мапе, добавляем его в результат
		if set2[elem] {
			result = append(result, elem)
		}
	}
	return result
}

// Метод с использованием срезов (Slices)
func intersectionWithSlice(set1, set2 []int) []int {
	// Сортируем срезы для удобства сравнения
	sort.Ints(set1)
	sort.Ints(set2)

	var result []int
	i, j := 0, 0

	// Проходим по срезам, сравнивая элементы
	for i < len(set1) && j < len(set2) {
		if set1[i] == set2[j] {
			// Если элементы равны, добавляем в результат и двигаемся к следующим элементам
			result = append(result, set1[i])
			i++
			j++
		} else if set1[i] < set2[j] {
			// Если элемент из первого множества меньше, увеличиваем индекс в первом множестве
			i++
		} else {
			// Если элемент из второго множества меньше, увеличиваем индекс во втором множестве
			j++
		}
	}
	return result
}

func main() {
	// Пример использования методов
	set1Map := []int{1, 2, 3, 4, 5}
	set2Map := []int{3, 4, 5, 6, 7}
	resultMap := intersectionWithMap(set1Map, set2Map)
	fmt.Println("Пересечение с помощью карт:", resultMap) // Вывод: map[3:true]

	set1Slice := []int{1, 2, 3, 4, 5}
	set2Slice := []int{3, 4, 5, 6, 7}
	resultSlice := intersectionWithSlice(set1Slice, set2Slice)
	fmt.Println("Пересечение методом слайсов:", resultSlice) // Вывод: [3]
}
