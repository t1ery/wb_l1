package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Слайс до удаления:", slice)
	i := 2 // Индекс элемента для удаления

	if i < len(slice) {
		slice = append(slice[:i], slice[i+1:]...)
		fmt.Println("Слайс после удаления:", slice)
	} else {
		fmt.Println("Неверный индекс элемента.")
	}
}
