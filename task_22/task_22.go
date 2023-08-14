package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Создаем новый объект для чтения ввода пользователя.
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите два числа (разделите пробелом):")
		// Читаем строку из ввода.
		input, _ := reader.ReadString('\n')
		// Убираем символ переноса строки из конца строки.
		input = strings.Trim(input, "\n")
		// Разделяем введенную строку на два числа.
		numbers := strings.Split(input, " ")
		if len(numbers) != 2 {
			fmt.Println("Пожалуйста, введите два числа.")
			continue // Продолжаем цикл сначала.
		}

		// Создаем объекты big.Int для хранения больших чисел.
		var a, b big.Int
		// Преобразуем введенные строки в большие числа.
		_, errA := a.SetString(numbers[0], 10)
		_, errB := b.SetString(numbers[1], 10)
		// Проверяем, успешно ли прошло преобразование.
		if errA != true || errB != true {
			fmt.Println("Ошибка ввода чисел. Пожалуйста, введите целые числа.")
			continue // Продолжаем цикл сначала.
		}

		fmt.Println("Выберите операцию:")
		fmt.Println("1. Умножение")
		fmt.Println("2. Деление")
		fmt.Println("3. Сложение")
		fmt.Println("4. Вычитание")

		// Читаем ввод пользователя для выбора операции.
		operationInput, _ := reader.ReadString('\n')
		operationInput = strings.Trim(operationInput, "\n")
		// Преобразуем строку в число.
		operation, err := strconv.Atoi(operationInput)
		// Проверяем корректность выбора операции.
		if err != nil || operation < 1 || operation > 4 {
			fmt.Println("Неверный выбор операции.")
			continue // Продолжаем цикл сначала.
		}

		// Создаем объект для хранения результата операции.
		var result big.Int

		// Выбираем операцию в зависимости от ввода пользователя.
		switch operation {
		case 1:
			// Умножение.
			result.Mul(&a, &b)
		case 2:
			// Деление.
			result.Div(&a, &b)
		case 3:
			// Сложение.
			result.Add(&a, &b)
		case 4:
			// Вычитание.
			result.Sub(&a, &b)
		}

		// Выводим результат операции.
		fmt.Printf("Результат: %s\n", result.String())

		fmt.Println("Хотите выполнить еще операцию? (да/нет)")
		// Читаем ввод пользователя для продолжения или завершения программы.
		repeatInput, _ := reader.ReadString('\n')
		repeatInput = strings.Trim(repeatInput, "\n")
		if strings.ToLower(repeatInput) != "да" {
			fmt.Println("Программа завершена.")
			break // Завершаем цикл.
		}
	}
}
