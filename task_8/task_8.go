package main

import (
	"fmt"
)

// Функция setBitwiseOps реализует установку i-го бита в числе используя побитовые операции.
func setBitwiseOps(num int64, pos uint, bitValue int) int64 {
	// Создаем маску, устанавливая бит в позиции pos (остальные биты сброшены).
	mask := int64(1) << pos

	// Если bitValue равен 0, инвертируем1 маску, чтобы сбросить бит в позиции pos.
	if bitValue == 0 {
		mask = ^mask
		return num & mask
	}

	// Иначе, используем побитовое ИЛИ, чтобы установить бит в позиции pos.
	return num | mask
}

// Функция setBitShift реализует установку i-го бита в числе используя побитовый сдвиг.
func setBitShift(num int64, pos uint, bitValue int) int64 {
	// Используем побитовый сдвиг для создания маски с установленным битом в позиции pos.
	// Затем, в зависимости от bitValue, либо устанавливаем бит, либо сбрасываем его.
	if bitValue == 1 {
		num |= 1 << pos
	} else if bitValue == 0 {
		num &= ^(1 << pos)
	}
	return num
}

func main() {
	var num int64
	fmt.Print("Введите число (int64): ")
	_, _ = fmt.Scan(&num)

	var pos uint
	fmt.Print("Введите позицию бита (0-63): ")
	_, _ = fmt.Scan(&pos)

	var bitValue int
	fmt.Print("Введите значение бита (0 или 1): ")
	_, _ = fmt.Scan(&bitValue)

	// Проверяем корректность введенной позиции бита.
	if pos < 0 || pos > 63 {
		fmt.Println("Неверная позиция бита.")
		return
	}

	// Проверяем корректность введенного значения бита.
	if bitValue != 0 && bitValue != 1 {
		fmt.Println("Неверное значение бита.")
		return
	}

	// Используем обе функции для установки бита и выводим результаты.
	numWithBitwiseOps := setBitwiseOps(num, pos, bitValue)
	fmt.Printf("Число с установленным битом (побитовые операции): %d\n", numWithBitwiseOps)

	numWithBitShift := setBitShift(num, pos, bitValue)
	fmt.Printf("Число с установленным битом (побитовый сдвиг): %d\n", numWithBitShift)
}
