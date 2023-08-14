package main

import (
	"fmt"
	"strings"
)

// reverseStringSlice переворачивает строку, используя срез рун.
func reverseStringSlice(input string) string {
	runes := []rune(input) // Преобразование строки в срез рун
	length := len(runes)

	// Переворачивание строки в срезе рун
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Меняем местами руны
	}

	return string(runes) // Преобразование среза рун обратно в строку
}

// reverseStringBuilder переворачивает строку, используя strings.Builder.
func reverseStringBuilder(input string) string {
	var reversed strings.Builder

	runes := []rune(input) // Преобразование строки в срез рун
	length := len(runes)

	// Построение строки с обратным порядком символов
	for i := length - 1; i >= 0; i-- {
		reversed.WriteRune(runes[i]) // Записываем руну в обратном порядке
	}

	return reversed.String() // Получение строки из strings.Builder
}

// reverseStringBytes переворачивает строку, используя байтовый массив.
func reverseStringBytes(input string) string {
	runes := []rune(input) // Преобразование строки в срез рун
	length := len(runes)

	reversed := make([]byte, 0, len(input)) // Создание байтового среза для переворота

	// Запись байтов в обратном порядке
	for i := length - 1; i >= 0; i-- {
		reversed = append(reversed, []byte(string(runes[i]))...) // Преобразование руны в байты и добавление к срезу
	}

	return string(reversed) // Преобразование байтового среза в строку
}

func main() {
	input := "?ил кат ен ,оньлокирП"

	// Переворачивание строки с использованием среза рун
	reversedSlice := reverseStringSlice(input)
	fmt.Println("Результат с использованием среза рун:", reversedSlice)

	// Переворачивание строки с использованием strings.Builder
	reversedBuilder := reverseStringBuilder(input)
	fmt.Println("Результат с использованием strings.Builder:", reversedBuilder)

	// Переворачивание строки с использованием байтового массива
	reversedBytes := reverseStringBytes(input)
	fmt.Println("Результат с использованием байтового массива:", reversedBytes)
}
