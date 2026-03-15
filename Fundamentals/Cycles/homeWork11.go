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
}
