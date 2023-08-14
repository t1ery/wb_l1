package main

import (
	"fmt"
	"time"
)

const N = 5 // Количество секунд

func sender(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond) // Отправляем значение в канал каждые 500 миллисекунд
	}
}

func receiver(ch <-chan int, doneChan <-chan struct{}) {
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				// Канал закрыт, завершаем работу
				return
			}
			fmt.Printf("Полученное значение: %d\n", value)
		case <-doneChan:
			// Получен сигнал завершения, завершаем работу
			return
		}
	}
}

func main() {
	ch := make(chan int)                       // Канал для передачи данных
	doneChan := make(chan struct{})            // Канал для сигнализации завершения работы
	timeoutChan := time.After(N * time.Second) // Таймер с кейсом времени N секунд

	// Запускаем горутину для отправки значений в канал
	go sender(ch)

	// Запускаем горутину для чтения из канала
	go receiver(ch, doneChan)

	// Ожидаем сигнала завершения или истечения времени
	select {
	case <-timeoutChan:
		// Истекло время выполнения, завершаем программу
		fmt.Println("Время выполнения истекло.")
	case <-doneChan:
		// Получен сигнал завершения, завершаем программу
	}

	// Закрываем канал сигнала завершения, чтобы завершить горутины sender и receiver
	close(doneChan)

	// Ждем некоторое время, чтобы дать горутинам sender и receiver возможность завершиться
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Программа завершена.")
}
