package main

import (
	"fmt"
)

func main() {
	//zadanie 1
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	// zadanie 2

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//zadanie 3
	number := 9
	for i := 1; i <= 10; i++ {
		fmt.Println(number, " x", i, "=", number*i)
	}
	//zadanie 4
	var n int
	fmt.Println("Введите число:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println(i)

		}
	}
	//zadanie 5
	var number_1 int
	fmt.Println("Введите число:")
	fmt.Scan(&number_1)
	score := 0

	for number_1 != 0 {
		number_1 = number_1 / 10
		score++
	}

	fmt.Println("Количество цифр в числе:", score)

	//zadanie 6
	text := "Developer"

	for _, v := range text {
		fmt.Printf("%c\n", v)
	}

	//zadanie 7

	balance := 3_000

	for {
		var n int
		fmt.Println("Введите число:")
		fmt.Scan(&n)
		if n == 1 {
			fmt.Println("Balance:", balance)

		} else if n == 2 {
			balance += 500
		} else if n == 3 {
			balance -= 200
		} else if n == 0 {
			fmt.Println("Выход из программы")
			break
		}
	}
	fmt.Println("Цикл завершен...")

	// zadanie 8
	// Пользователь вводит число n. Используя один цикл for, выведите последовательность чисел от 1 до n.
	// При этом программа должна вести счёт:
	// сколько чисел меньше 10,
	// сколько чисел от 10 до 20 включительно,
	// и сколько чисел больше 20.

	// После завершения цикла выведите три результата: количество чисел в каждом из этих диапазонов.

	var n_1 int
	fmt.Println("Введите число:")
	fmt.Scan(&n_1)
	count_10 := 0
	count_10_20 := 0
	count := 0

	for i := 1; i <= n_1; i++ {
		if i < 10 {
			count_10++
		} else if i >= 10 && i <= 20 {
			count_10_20 += 1
		} else if i > 20 {
			count += 1
		}

	}
	fmt.Println("Чисел от 0 до 10:", count_10)
	fmt.Println("Чисел от 10 до 20:", count_10_20)
	fmt.Println("Чисел больше 20:", count)

	// // zadanie 9
	// // Используя вложенные циклы for, выведите треугольник из символов *.
	// // Создайте переменную n, которая задаёт количество строк.
	// // В первой строке должен быть один *, во второй — два, в третьей — три и так далее, пока количество * не станет равно n

	var star_1 int
	fmt.Println("Введите число:")
	fmt.Scan(&star_1)

	for i := 1; i <= star_1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
