package main

import (
	"fmt"
	"strings"
)

func main() {
	//Zadanie 1
	temperature := -27
	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature <= 20 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}
	//zadanie 2
	score := 89
	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 70 { //&& score <= 89 {
		fmt.Println("Хорошо")
	} else if score >= 50 { //&& score <= 69 {
		fmt.Println("Хорошо")
	} else {
		fmt.Println("Жарко")
	}

	//Zadanie 3
	hours := 4
	switch hours {
	case 0, 1, 2, 3, 4, 5:
		fmt.Println("Ночь")
	case 6, 7, 8, 9, 10, 11:
		fmt.Println("Утро")
	case 12, 13, 14, 15, 16, 17:
		fmt.Println("День")
	case 18, 19, 20, 21, 22, 23:
		fmt.Println("Вечер")
	}
	//zadanie 4
	var number int
	fmt.Print("Введите номер: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("Чётное число")
	} else {
		fmt.Println("Нечётное число")
	}

	//zadanie 5
	var day string
	fmt.Print("Введите день недели:")
	fmt.Scan(&day)

	switch strings.ToLower(day) {
	case "monday", "muesday", "wednesday", "thursday", "friday":
		fmt.Println("Будний день")
	case "saturday", "sunday":
		fmt.Println("Выходной день")
	default:
		fmt.Println("Некорректный день")
	}

	//zadanie 6
	balance := 9_000_000
	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}

	//zadanie 7
	var age int
	fmt.Print("Введите Ваш возраст:")
	fmt.Scan(&age)
	if age < 13 {
		fmt.Println("Ребенок")
	} else if age <= 17 {
		fmt.Println("Подросток")
	} else if age <= 59 {
		fmt.Println("Взрослый")
	} else if age >= 60 {
		fmt.Println("Пожилой")
	}
	//zadanie 8
	var command string
	fmt.Println("Введите команду")
	fmt.Scan(&command)

	switch strings.ToLower(command) {
	case "start":
		fmt.Println("start")
	case "stop":
		fmt.Println("stop")
	case "restart":
		fmt.Println("restart")
	default:
		fmt.Println("Неизвестная команда")
	}

	//Zadanie 9
	var grade float64
	fmt.Print("Оценка:")
	fmt.Scan(&grade)
	switch {
	case grade >= 5.3:
		fmt.Println("A++")
	case grade > 5.0:
		fmt.Println("A+")
	case grade == 5:
		fmt.Println("A")
	case grade >= 4.5:
		fmt.Println("A-")

	case grade > 4.0:
		fmt.Println("B+")
	case grade == 4:
		fmt.Println("B")
	case grade >= 3.5:
		fmt.Println("B-")

	case grade > 3:
		fmt.Println("C+")
	case grade == 3:
		fmt.Println("C")
	case grade >= 2.5:
		fmt.Println("C-")

	case grade > 2:
		fmt.Println("D+")
	case grade == 2:
		fmt.Println("D")
	case grade >= 1.5:
		fmt.Println("D-")

	case grade == 1:
		fmt.Println("F")
	default:
		fmt.Println("Некоректная оценка")
	}

}
