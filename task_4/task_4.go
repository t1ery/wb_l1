package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const N = 5 // Количество воркеров

func worker(id int, dataChan <-chan int, wg *sync.WaitGroup, doneChan <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				// Канал dataChan закрыт, завершаем работу
				return
			}
			time.Sleep(1500 * time.Millisecond) // Задержка в 1.5 секунды
			fmt.Printf("Воркер %d: Полученные данные %d\n", id, data)
		case <-doneChan:
			// Получен сигнал завершения, завершаем работу
			return
		}
	}
}

func main() {
	dataChan := make(chan int) // Канал для передачи данных от главного потока к воркерам

	var wg sync.WaitGroup
	wg.Add(N)

	// Создаем канал для сигнализации завершения работы воркерам
	doneChan := make(chan struct{})

	// Запускаем N воркеров
	for i := 1; i <= N; i++ {
		go worker(i, dataChan, &wg, doneChan)
	}

	// Запись данных в канал из главного потока
	go func() {
		for i := 1; ; i++ {
			dataChan <- i
		}
	}()

	// Ожидаем сигнала завершения (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Ожидание сигнала завершения
	<-sigChan

	// Закрываем канал, чтобы сообщить воркерам о завершении работы
	close(doneChan)

	// Дожидаемся завершения работы всех воркеров
	wg.Wait()

	// Закрываем канал, чтобы завершить операции чтения из dataChan
	close(dataChan)

	fmt.Println("Программа завершена.")
}
