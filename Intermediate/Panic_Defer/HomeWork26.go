package main

import (
	"fmt"
)

// zadanie 1
func printData() {
	fmt.Println("Обработка данных")
}

func checkAge(age int) int {
	if age < 0 {
		panic("Возрат не может быть меньше отрицательным")
	}
	return 0
}

// zadanie 4
func cleanup() {
	fmt.Println("Очистка ресурсов выполнена")
}

// zadanie 5
func handlePanic() {

	if r5 := recover(); r5 != nil {
		fmt.Println("Паника успешно перехвачена:", r5)
	}

}

// zadanie 6
func safeDivide(num1, num2 float64) {
	defer func() {
		if r6 := recover(); r6 != nil {
			fmt.Println("Паника успешно перехвачена:", r6)
		}
	}()
	if num2 == 0 {
		panic("Деление на ноль!")

	}
	fmt.Println("Результат:", num1/num2)

}

func getElement(num []int, index int) int {
	defer func() {
		if r7 := recover(); r7 != nil {
			fmt.Println("Паника успешно перехвачена:", r7)
		}
	}()
	return num[index]
}

func main() {
	// //zadanie 1
	// fmt.Println("Старт программы")
	// defer printData()
	// fmt.Println("КОнец программы")

	// //zadanie 2
	// defer printData()
	// fmt.Println("Первый")
	// defer printData()

	// fmt.Println("Второй")
	// defer printData()
	// fmt.Println("Третий")

	// //zadanie 3
	// // fmt.Println(checkAge(-5))
	// // fmt.Println("КОнец программы")

	// //zadanie 4
	// defer cleanup()
	// panic("Критическая ошибка!")

	//zadanie 5
	// defer handlePanic()

	// fmt.Println("Старт программы")
	// panic("чек код!")
	// fmt.Println("КОнец программы") // не выполнится

	//zadanie 6
	safeDivide(52, 2)
	safeDivide(52, 0)

	//zadanie 7
	num7 := []int{1, 2, 3, 45, 7, 6}

	fmt.Println(getElement(num7, 5))
	fmt.Println(getElement(num7, 25))

}
