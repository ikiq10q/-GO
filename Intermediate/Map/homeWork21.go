package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// //zadanie 1
	// //Вы создаёте первую программу на Go после установки языка. Создайте map toolUsage, где
	// //ключ - название инструмента ("Go", "VSCode", "Git"), а значение - количество раз, которое вы использовали этот инструмент за день.
	// //Присвойте значения 3, 5 и 2 соответственно. Используя цикл for range, выведите название каждого инструмента и количество использований.

	toolUsage := map[string]int{"Go": 3, "VSCode": 5, "Git": 2}

	for key, values := range toolUsage {
		fmt.Println(key, values)
	}

	//zadanie 2
	//В программе хранится информация о запуске Go-приложения. Создайте map buildStatus, где ключ - строка "build" или "run",
	// а значение - true или false.Установите "build" в true, а "run" в false.
	// Если "build" равен true, выведите сообщение "Сборка прошла успешно".
	buildStatus := map[string]bool{
		"build": true,
		"run":   false,
	}

	if buildStatus["build"] {
		fmt.Println("Сборка прошла успешно")
	}

	// zadanie 3
	// Создайте программу, которая запрашивает у пользователя его имя.
	// Создайте map userInfo, где ключ "name" хранит введённое имя, а ключ "isLoggedIn" хранит значение "true".
	// Выведите сообщение "Пользователь <имя> вошёл в систему".

	var name string

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	userInfo := map[string]string{
		"name":       name,
		"isLoggedIn": "true",
	}

	if userInfo["isLoggedIn"] == "true" {
		fmt.Printf("Пользователь %s вошёл в систему\n", userInfo["name"])
	}

	//zadanie 4
	//Создайте map cpuLoad, где ключ - номер ядра процессора (int), а значение - загрузка в процентах (int).
	// Добавьте данные для ядер 1, 2 и 3 со значениями 40, 65 и 30.
	// Найдите и выведите номер ядра с максимальной загрузкой, используя цикл.
	cpuLoad := map[int]int{
		1: 40,
		2: 65,
		3: 30,
	}

	maxVal := math.MinInt
	var maxCpu int
	fl := false

	for num, cpu := range cpuLoad {
		if cpu > maxVal {
			maxVal = cpu
			maxCpu = num
			fl = true
		}
	}

	if fl {
		fmt.Printf("Максимальная загрузка: %d, найдено у ядра: %d\n", maxVal, maxCpu)

	}

	//zadanie 5
	//В программе хранятся оценки студентов. Создайте map examResults, где ключ - имя студента, а значение - оценка.
	// Добавьте "Aruzhan" = 85, "Dias" = 92, "Alina" = 78.
	// Используя условия, выведите имя каждого студента и сообщение "<имя> сдал экзамен", если оценка больше или равна 80, иначе "<имя> не сдал экзамен".

	examResults := map[string]int{
		"Aruzhan": 85,
		"Dias":    92,
		"Alina":   78,
	}
	for students, points := range examResults {
		if points >= 80 {
			fmt.Printf("%s сдал экзамен\n", students)
		} else {
			fmt.Printf("%s не сдал экзамен\n", students)
		}
	}

	//zadanie 6
	//Создайте слайс words со значениями "go", "is", "fast".
	//Создайте map wordLength, где ключ - слово, а значение - длина этого слова.
	//Используя цикл, заполните map и выведите его содержимое.

	words := []string{"go", "is", "fast"}
	wordLength := make(map[string]int)

	for _, word := range words {
		wordLength[word] = len(word)
	}

	fmt.Println(wordLength)

	//zadanie 7
	//Создайте map menu, где ключ - название блюда, а значение - цена.
	//Добавьте "Burger" = 2500, "Pizza" = 3200, "Tea" = 500. Запросите у пользователя название блюда.
	//Если блюдо есть в map, выведите его цену, иначе выведите сообщение "Блюдо не найдено".

	menu := map[string]int{
		"burger": 2500,
		"pizza":  3200,
		"tea":    500,
	}
	var zapros string
	fmt.Print("Выберете из меню:")
	fmt.Scan(&zapros)

	zapros = strings.ToLower(zapros)

	price, key := menu[zapros]

	if key {
		fmt.Printf("Цена: %d\n", price)
	} else {
		fmt.Println("Блюдо не найдено")
	}

	//zadanie 8
	//Создайте map loginAttempts, где ключ - логин пользователя, а значение - количество попыток входа.
	//Добавьте "admin" = 2 и "guest" = 0. Увеличьте количество попыток "admin" на 1.
	//Если количество попыток для "admin" стало больше 2, выведите сообщение "Доступ заблокирован".
	loginAttempts := map[string]int{
		"admin": 2,
		"guest": 0,
	}

	loginAttempts["admin"]++

	if loginAttempts["admin"] > 2 {
		fmt.Println("Доступ заблокирован")
	}

	//zadanie 9
	//Создайте двумерный массив sales, который хранит продажи за 3 дня для 2 магазинов.
	//Количество продаж заполните самостоятельно. Создайте пустой map total с где ключ - номер магазина, значение - общее количество продаж.
	//Используя вложенный цикл for заполните map и выведите его в консоль.
	sales := [2][3]int{
		{10, 20, 15},
		{5, 7, 12},
	}

	total := make(map[int]int)

	for i := 0; i < len(sales); i++ {
		sum := 0
		for j := 0; j < len(sales[i]); j++ {
			sum += sales[i][j]
		}
		total[i] = sum
	}
	fmt.Println(total)

	//zadanie 10
	//Создайте слайс numbers со значениями .
	//Создайте map numberStatus, где ключ - число, а значение - строка "even" или "odd".
	//Используя цикл и условие четности чисел, заполните map и выведите результат.

	numbers := []int{4, 7, 2, 9, 5}
	numberStatus := make(map[int]string)

	for _, num := range numbers {
		if num%2 == 0 {
			numberStatus[num] = "even"
		} else {
			numberStatus[num] = "odd"
		}
	}

	fmt.Println(numberStatus)

	//zadanie 11
	//Вы разрабатываете систему проверки конфигурации сервера.
	// У вас есть «эталонные» настройки и «текущие» настройки сервера. Необходимо проверить, совпадают ли они полностью.
	//
	//Создайте две map. Первая map defaultConfig должна содержать: "host" = "localhost", "port" = "8080", "mode" = "production"
	//Вторая map currentConfig должна содержать те же значения.
	//Сравните эти две map, если они полностью совпадают - выведите сообщение  "Конфигурации совпадают".
	//Если нет - выведите "Конфигурации отличаются". После этого измените значение "mode" во второй map на "debug" и снова выполните сравнение. Выведите результат.

	defaultConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	currentConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}
	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

	currentConfig["mode"] = "debug"

	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

}
