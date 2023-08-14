package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]int)
	var mutex sync.Mutex     // Mutex для синхронизации записи
	var rwMutex sync.RWMutex // RWMutex для синхронизации чтения и записи
	var wg sync.WaitGroup

	numRoutines := 10
	numOperations := 100

	// Способ 1: Mutex
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				mutex.Lock()
				data[id] = j
				mutex.Unlock()
			}
		}(i)
	}
	wg.Wait()

	// Вывод результатов
	fmt.Println("Результат работы 1-го способа:")
	for key, value := range data {
		fmt.Printf("%d: %d\n", key, value)
	}

	// Очистка данных
	data = make(map[int]int)

	// Способ 2: RWMutex (чтение)
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				rwMutex.RLock()
				_ = data[id]
				rwMutex.RUnlock()
			}
		}(i)
	}
	wg.Wait()

	// Вывод результатов
	fmt.Println("Результат работы 2-го способа(чтение):")
	for key, value := range data {
		fmt.Printf("%d: %d\n", key, value)
	}

	// Очистка данных
	data = make(map[int]int)

	// Способ 3: RWMutex (запись)
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				rwMutex.Lock()
				data[id] = j
				rwMutex.Unlock()
			}
		}(i)
	}
	wg.Wait()

	// Вывод результатов
	fmt.Println("Результат работы 3-го способа:")
	for key, value := range data {
		fmt.Printf("%d: %d\n", key, value)
	}
}
