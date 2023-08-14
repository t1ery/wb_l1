package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализация с мьютексами
type MutexCounter struct {
	value int
	mu    sync.Mutex
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *MutexCounter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Реализация с использованием sync/atomic
type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

// Реализация с использованием каналов
type ChannelCounter struct {
	value int
}

func Increment(counter *ChannelCounter, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	counter.value++
	ch <- counter.value
}

func main() {
	var wg sync.WaitGroup

	// Использование MutexCounter
	mutexCounter := MutexCounter{}
	numGoroutines := 10
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			mutexCounter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Значение счётчика MutexCounter:", mutexCounter.GetValue())

	// Использование AtomicCounter
	atomicCounter := AtomicCounter{}
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			atomicCounter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Значение счётчика AtomicCounter:", atomicCounter.GetValue())

	// Использование ChannelCounter
	channelCounter := ChannelCounter{}
	ch := make(chan int, numGoroutines) // Буферизованный канал
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go Increment(&channelCounter, &wg, ch)
	}
	wg.Wait()
	close(ch)

	// Собираем значения из канала
	finalValue := 0
	for value := range ch {
		finalValue = value
	}

	fmt.Println("Значение счётчика ChannelCounter:", finalValue)
}
