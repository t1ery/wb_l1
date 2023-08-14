package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userInput string

	fmt.Print("Введите значение: ")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	var value interface{}

	// Попытка преобразования в int
	if intValue, err := strconv.Atoi(userInput); err == nil {
		value = intValue
	} else {
		// Попытка преобразования в bool
		if boolValue, err := strconv.ParseBool(userInput); err == nil {
			value = boolValue
		} else {
			// Попытка создания chan
			if makeChan, ok := makeChanFromString(userInput); ok {
				value = makeChan
			} else {
				value = userInput
			}
		}
	}

	fmt.Printf("Вы ввели значение: %v\n", value)
	fmt.Printf("Тип введенного значения: %T\n", value)
}

// Вспомогательная функция для создания chan из строки
func makeChanFromString(input string) (chan int, bool) {
	if input == "chan" {
		return make(chan int), true
	}
	return nil, false
}
