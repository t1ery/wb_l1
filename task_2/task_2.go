package main

import (
	"fmt"
	"sync"
)

// Функция calculateSquare рассчитывает квадрат числа и отправляет его в канал result.
func calculateSquare(number int, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()
	square := number * number
	result <- square
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал result для обмена результатами между горутинами.
	result := make(chan int)

	// Инициализируем WaitGroup для синхронизации горутин.
	var wg sync.WaitGroup

	// Запускаем горутины для рассчета квадратов чисел.
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, &wg, result)
	}

	// Дополнительная горутина, которая будет ждать завершения всех горутин и закроет канал result.
	go func() {
		wg.Wait()
		close(result)
	}()

	// Главная горутина (функция main) читает результаты из канала result и выводит их в stdout.
	for square := range result {
		fmt.Println(square)
	}
}

// Главная горутина (функция main) создает отдельную горутину для каждого числа из массива,
// и каждая горутина рассчитывает квадрат числа асинхронно.
// Затем главная горутина читает результаты из канала result и выводит их в stdout.
// Когда все горутины завершат свою работу, главная горутина закрывает канал result,
// чтобы дать знать, что все результаты получены, и
// таким образом главная горутина выходит из цикла чтения из канала.
