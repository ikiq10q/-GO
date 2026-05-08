package main

import "fmt"

func main() {
	//zadanie 1
	result := square(4)
	fmt.Println(result)

	//zadanie 2
	fmt.Println(maxNumber(2, 5))

	//zadanie 3
	fmt.Println(isEven(2))
	fmt.Println(isEven(3))
	fmt.Println(isEven(6))

	//zadanie 4
	greetUser("Aselhan")
	greetUser("Daniil")

	//zadanie 5
	slce := []int{10, 20, 52, 55, 44}
	sum := sumSmice(slce)

	fmt.Printf("Сумма: %d\n", sum)

	//zadanie 6
	fmt.Println("Admin:", checkLogin("admin", "1234"))
	fmt.Println("Users:", checkLogin("ZXC", "qwerty"))

	//zadanie 7
	balance := 5000
	increaseBalance(&balance, 1000)
	fmt.Printf("Ваш баланс: %d\n", balance)

	//zadanie 8
	attempts := 5
	resetAttempts(&attempts)
	fmt.Printf("У Вас осталось попыток сброса: %d\n", attempts)

}

//zadanie1
func square(x int) (sqr int) {
	sqr = x * x
	return sqr
}

//zadanie2
func maxNumber(x, y int) (maxN int) {
	if x > y {
		return x
	} else {
		return y
	}
	// maxN = max(x, y)
	// return maxN
}

//zadanie 3

func isEven(x int) (isEven bool) {

	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

//zadanie 4

func greetUser(name string) {

	fmt.Printf("Привет, %s!\n", name)

}

//zadanie 5

func sumSmice(num []int) (sumSmice int) {
	score := 0
	for _, v := range num {
		score += v
	}
	return score
}

//zadanie 6
func checkLogin(login, password string) bool {
	admin := map[string]string{
		"admin": "1234",
	}

	pass, ok := admin[login]
	if pass == password && ok {
		return true
	} else {
		return false
	}
}

//zadanie 7
//Создайте функцию increaseBalance, которая принимает указатель на переменную типа int (баланс) и число типа int (сумма пополнения).
//Внутри функции увеличьте значение баланса на переданную сумму. В main создайте переменную balance, передайте её в функцию и выведите новое значение.
func increaseBalance(x *int, summ int) {
	*x += summ

}

//zadanie 8

func resetAttempts(attempts *int) {
	if *attempts > 3 {
		*attempts = 0
	}
}
