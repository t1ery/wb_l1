package main

import (
	"fmt"
	"strings"
)

// reverseWordsSlice переворачивает слова в строке, используя срез рун.
func reverseWordsSlice(input string) string {
	words := strings.Fields(input) // Разбиваем строку на слова

	// Переворачиваем порядок слов в срезе
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // Меняем местами слова
	}

	return strings.Join(words, " ") // Объединяем слова в строку с пробелами
}

// reverseWordsBuilder переворачивает слова в строке, используя strings.Builder.
func reverseWordsBuilder(input string) string {
	words := strings.Fields(input) // Разбиваем строку на слова

	var reversed strings.Builder

	// Построение строки с обратным порядком слов
	for i := len(words) - 1; i >= 0; i-- {
		reversed.WriteString(words[i]) // Добавляем слово в strings.Builder
		if i > 0 {
			reversed.WriteRune(' ') // Добавляем пробел, кроме последнего слова
		}
	}

	return reversed.String() // Получение строки из strings.Builder
}

// reverseWordsBytes переворачивает слова в строке, используя байтовый срез.
func reverseWordsBytes(input string) string {
	words := strings.Fields(input) // Разбиваем строку на слова

	var reversed []byte

	// Запись байтов в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, []byte(words[i])...)
		if i > 0 {
			reversed = append(reversed, ' ')
		}
	}

	return string(reversed) // Преобразование байтового среза в строку
}

func main() {
	input := "Какая то абракадабра"

	// Переворот слов с использованием среза рун
	reversedSlice := reverseWordsSlice(input)
	fmt.Println("Результат с использованием среза рун:", reversedSlice)

	// Переворот слов с использованием strings.Builder
	reversedBuilder := reverseWordsBuilder(input)
	fmt.Println("Результат с использованием strings.Builder:", reversedBuilder)

	// Переворот слов с использованием байтового среза
	reversedBytes := reverseWordsBytes(input)
	fmt.Println("Результат с использованием байтового среза:", reversedBytes)
}
