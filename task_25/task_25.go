package main

import (
	"fmt"
	"time"
)

// SleepUsingTime - реализация с помощью time.Sleep
func SleepUsingTime(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// SleepUsingTimer - реализация с помощью таймера
func SleepUsingTimer(seconds int) {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer.C
}

// SleepUsingGoroutine - реализация с помощью горутины и канала
func SleepUsingGoroutine(seconds int, done chan bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	done <- true
}

// SleepUsingSelect - реализация с помощью select
func SleepUsingSelect(seconds int) {
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		return
	}
}

func main() {
	fmt.Println("В самом начале")

	// Используя time.Sleep
	SleepUsingTime(2)
	fmt.Println("После time.Sleep")

	// Используя таймер
	SleepUsingTimer(3)
	fmt.Println("После таймера")

	// Используя горутину и канал
	done := make(chan bool)
	go SleepUsingGoroutine(4, done)
	<-done
	fmt.Println("После горутины")

	// Используя select
	SleepUsingSelect(5)
	fmt.Println("После select")
}
