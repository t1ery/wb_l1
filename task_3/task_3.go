package main

import (
	"fmt"
	"sync"
)

// Функция calculateSquare вычисляет квадрат числа и отправляет его в канал result.
func calculateSquare(number int, result chan int) {
	square := number * number
	result <- square
}

// Функция calculateSquareWithMutex вычисляет квадрат числа и добавляет его к общей сумме.
// Использует sync.Mutex для защиты общей переменной totalSum от гонки данных при параллельном доступе.
func calculateSquareWithMutex(number int, mu *sync.Mutex, totalSum *int) {
	square := number * number
	mu.Lock()
	*totalSum += square
	mu.Unlock()
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// === Первое решение: с использованием каналов ===

	// Создаем буферизованный канал для передачи результатов вычислений.
	resultCh := make(chan int, len(numbers))
	// WaitGroup для ожидания завершения горутин вычисления квадратов.
	var wgCh sync.WaitGroup

	for _, num := range numbers {
		wgCh.Add(1)
		go func(n int) {
			defer wgCh.Done()
			calculateSquare(n, resultCh)
		}(num)
	}

	// Запускаем горутину для ожидания завершения всех горутин вычисления квадратов.
	go func() {
		wgCh.Wait()
		close(resultCh)
	}()

	var sumCh int
	// Считываем результаты из канала и суммируем квадраты.
	for square := range resultCh {
		sumCh += square
	}

	fmt.Println("Сумма квадратов (с использованием каналов):", sumCh)

	// === Второе решение: с использованием Mutex ===

	var totalSumMutex int
	// WaitGroup для ожидания завершения горутин вычисления квадратов.
	var wgMutex sync.WaitGroup
	// Mutex для защиты общей переменной totalSumMutex от гонки данных при параллельном доступе.
	var mu sync.Mutex

	for _, num := range numbers {
		wgMutex.Add(1)
		go func(n int) {
			defer wgMutex.Done()
			calculateSquareWithMutex(n, &mu, &totalSumMutex)
		}(num)
	}

	// Ожидаем завершения всех горутин.
	wgMutex.Wait()

	fmt.Println("Сумма квадратов (с использованием Mutex):", totalSumMutex)
}
