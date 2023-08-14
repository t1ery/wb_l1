package main

import (
	"fmt"
	"math"
)

func main() {
	// Исходные температуры
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем словарь для группировки температур по диапазонам
	groupedTemperatures := make(map[int][]float64)

	// Перебираем исходные температуры
	for _, temp := range temperatures {
		// Определяем номер группы для текущей температуры
		group := int(math.Round(temp/10.0)) * 10

		// Добавляем текущую температуру в соответствующую группу
		groupedTemperatures[group] = append(groupedTemperatures[group], temp)
	}

	// Выводим группы температур
	for group, temps := range groupedTemperatures {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
