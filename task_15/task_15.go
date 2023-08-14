package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	// Создание огромной строки
	v := createHugeString(1 << 10)

	// Выделение памяти для нового среза размером 100 байт
	justStringBytes := make([]byte, 100)

	// Копирование первых 100 байт из огромной строки в новый срез
	copy(justStringBytes, v[:100])

	// Преобразование среза байтов обратно в строку
	justString = string(justStringBytes)
}

func main() {
	someFunc()

	// Вывод строки justString
	fmt.Println(justString)
}

// Функция для создания огромной строки
func createHugeString(size int) string {
	// Создание строки, повторяющейся множество раз, чтобы получить огромный размер
	return strings.Repeat("X", size)
}

// Данный фрагмент кода (который был в задании) на языке Go может привести к утечке памяти из-за использования среза v[:100]
// для присваивания значения переменной justString. При этом будет сохранена ссылка на огромную строку в памяти,
// и даже если она больше не используется, она не будет освобождена сборщиком мусора,
// так как ссылка на неё сохранена в justString.
