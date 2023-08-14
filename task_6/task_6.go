package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var stopFlag bool // Флаг для остановки горутин

// Пример использования каналов для синхронизации горутин
func usingChannels() {
	fmt.Println("Использование каналов:")

	done := make(chan struct{}) // Канал для сигнала завершения горутины

	go func() {
		defer fmt.Println("Горутина на каналах: Завершено!")
		defer close(done)
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина на каналах:", i)
			time.Sleep(time.Second)
		}
	}()

	// Ждем сигнала завершения
	<-done
}

// Пример использования контекста для остановки горутины
func usingContext() {
	fmt.Println("Использование контекста:")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer fmt.Println("Горутина на контексте: Завершено!")
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина на контексте:", i)
				time.Sleep(time.Second)
			}
		}
	}()

	// Ждем сигнала завершения (нажатия клавиши)
	fmt.Println("Нажмите Enter, чтобы остановить горутину на контексте.")
	fmt.Scanln()
	cancel()
}

// Пример использования WaitGroup для синхронизации горутин
func usingWaitGroup() {
	fmt.Println("Использование WaitGroup:")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer fmt.Println("Горутина на WaitGroup: Завершено!")
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина на WaitGroup:", i)
			time.Sleep(time.Second)
		}
	}()

	// Ждем завершения работы горутины
	wg.Wait()
}

// Пример использования сигналов для остановки горутин
func usingSignals() {
	fmt.Println("Использование сигналов:")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	done := make(chan struct{})

	go func() {
		defer fmt.Println("Горутина на сигналах: Завершено!")
		<-sigChan
		close(done)
	}()

	// Ожидание сигнала завершения (Ctrl+C)
	fmt.Println("Нажмите Ctrl+C, чтобы остановить горутину на сигналах.")
	<-done
}

// Пример использования runtime.Goexit для завершения горутины
func usingGoexit() {
	fmt.Println("Использование runtime.Goexit:")

	go func() {
		defer fmt.Println("Горутина на runtime.Goexit: Завершено!")
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина на runtime.Goexit:", i)
			time.Sleep(time.Second)
			if i == 2 {
				runtime.Goexit() // Завершение текущей горутины
			}
		}
	}()

	// Ожидание для завершения вывода перед завершением программы
	time.Sleep(6 * time.Second)
}

// Пример использования остановки горутины с помощью флага
func usingFlag() {
	fmt.Println("Использование флага:")

	stopFlag = false

	go func() {
		defer fmt.Println("Горутина с флагом: Завершено!")
		for i := 0; ; i++ {
			if stopFlag {
				return
			}
			fmt.Println("Горутина с флагом:", i)
			time.Sleep(time.Second)
		}
	}()

	// Ожидание нажатия клавиши для остановки
	fmt.Println("Нажмите Enter, чтобы остановить горутину с флагом.")
	fmt.Scanln()
	stopFlag = true
}

// Реализация остановки горутины с помощью паники (такое себе)
func usingPanic() {
	fmt.Println("Использование паники:")

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Паника перехвачена:", r)
			}
		}()
		defer fmt.Println("Горутина с паникой: Завершено!")
		for i := 0; ; i++ {
			fmt.Println("Горутина с паникой:", i)
			if i == 3 {
				panic("Пример паники") // Имитация паники
			}
			time.Sleep(time.Second)
		}
	}()

	// Ожидание для завершения вывода перед завершением программы
	time.Sleep(6 * time.Second)
}

func main() {
	usingChannels()
	fmt.Println()

	usingContext()
	fmt.Println()

	usingWaitGroup()
	fmt.Println()

	usingSignals()
	fmt.Println()

	usingFlag()
	fmt.Println()

	usingPanic()
	fmt.Println()

	usingGoexit()
}
