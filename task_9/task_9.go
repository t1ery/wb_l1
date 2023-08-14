package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4, 5}

	inputChannel := make(chan int)
	outputChannel := make(chan int)
	var wg sync.WaitGroup

	// Горутина для записи чисел в inputChannel
	go func() {
		for _, num := range input {
			inputChannel <- num
		}
		close(inputChannel)
	}()

	// Горутина для умножения чисел из inputChannel и отправки результата в outputChannel
	go func() {
		for num := range inputChannel {
			outputChannel <- num * 2
		}
		close(outputChannel)
	}()

	// Горутина для вывода данных из outputChannel в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range outputChannel {
			time.Sleep(1500 * time.Millisecond)
			fmt.Println(num)
		}
	}()

	// Ожидание завершения всех горутин
	wg.Wait()

	fmt.Println("Программа завершена")
}
