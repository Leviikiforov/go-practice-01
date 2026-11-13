package main

import "fmt"

func main() {
	var name string = "Георгий"
	group := "Учебная-01"
	const course = "Бэкенд-разработка на языке Go"
	complete := 3
	var interes bool = true

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Println("Курс:", course)
	fmt.Println("Завершено работ:", complete)
	fmt.Printf("Интерес к бэкенду: %t\n", interes)

}
