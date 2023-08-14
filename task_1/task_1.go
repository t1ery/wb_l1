package main

import "fmt"

// Human Определяем структуру
type Human struct {
	name   string
	age    int
	gender string
}

// Introduce Метод для структуры Human
func (h Human) Introduce() {
	fmt.Printf("Привет, меня зовут %s. Мне %d лет. Я %s.\n", h.name, h.age, h.gender)
}

// Action Определяем структуру с встраиванием структуры Human
type Action struct {
	Human
	hobby string
}

// DescribeHobby Метод для структуры Action
func (a Action) DescribeHobby() {
	fmt.Printf("Моё хобби - %s.", a.hobby)
}

func main() {
	// Создаем объект структуры Action
	person := Action{
		Human: Human{
			name:   "Артём",
			age:    25,
			gender: "Мужчина",
		},
		hobby: "Дайвинг",
	}

	// Теперь объект person имеет доступ к методам из структуры Human
	person.Introduce()
	person.DescribeHobby()

}
