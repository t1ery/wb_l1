package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция для проверки уникальности символов в строке
func areAllCharactersUnique(input string) bool {
	// Приведение строки к нижнему регистру для регистронезависимой проверки
	lowercaseInput := strings.ToLower(input)
	// Создание мапы для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	// Проход по каждому символу в строке
	for _, char := range lowercaseInput {
		// Если символ уже есть в мапе, то символ не уникален
		if charMap[char] {
			return false
		}
		// Добавление символа в мапу
		charMap[char] = true
	}

	return true
}

func main() {
	fmt.Println("Введите строки для проверки уникальности символов (для завершения введите 'exit'):")

	// Создание сканнера для чтения ввода пользователя
	scanner := bufio.NewScanner(os.Stdin)
	// Бесконечный цикл для ввода строк
	for scanner.Scan() {
		// Считывание введенной строки
		input := scanner.Text()
		// Проверка на завершение программы
		if input == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		// Вызов функции для проверки уникальности символов
		result := areAllCharactersUnique(input)
		// Вывод результата проверки
		fmt.Printf("%s — %v\n", input, result)
	}

	// Проверка на ошибки сканнера
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка ввода:", err)
	}
}
