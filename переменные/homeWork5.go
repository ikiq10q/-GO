package main

import "fmt"

func main() {
	// zadanie 1
	var schooling int
	schooling = 11
	fmt.Println("Мое школьное образование составляет:", schooling, "лет")
	schooling += 1
	fmt.Println("Мое школьное образование составляет:", schooling, "лет")

	// zadanie 2

	var name string
	name = "Vladislav"
	fmt.Println("Name:", name)
	name = "Ardak"
	fmt.Println("Name:", name)

	//zadanie 3
	steps := 0
	fmt.Println("Steps:", steps)
	steps += 2000
	fmt.Println("Steps:", steps)
	fmt.Println("Хорошая работа! Вы уже на пути к своей ежедневной цели!")

	//zadanie 4
	largeNumber := 9_000_000
	fmt.Println("LargeNumber:", largeNumber)

	//Zadanie 5
	// переменные тима const нельзя изменить, так как значение не изменяемо
	const breakTime = 15
	fmt.Println("BreakTime", breakTime)
	breakTime = 20
	fmt.Println("BreakTime", breakTime)
}
