package main

import (
	"fmt"
	"slices"
)

func main() {

	// //zadadnie 1

	// var numbers []int
	// num := 0

	// for {
	// 	fmt.Print("Введите число: ")
	// 	fmt.Scan(&num)

	// 	if num == 0 {
	// 		break
	// 	}
	// 	numbers = append(numbers, num)
	// }
	// fmt.Println(numbers)
	// summ := 0
	// for _, num := range numbers {
	// 	summ += num
	// }
	// fmt.Println(summ)

	// //zadanie 2
	// var values []int
	// num_1 := 0

	// for {
	// 	fmt.Print("Введите число: ")
	// 	fmt.Scan(&num_1)

	// 	if num_1 == 0 {
	// 		break
	// 	}
	// 	values = append(values, num_1)
	// }
	// fmt.Println("первый срез", values)

	// var values_1 []int

	// for _, v := range values {
	// 	if v%2 == 0 {
	// 		values_1 = append(values_1, v)
	// 	}
	// }

	// fmt.Println("Чётные числа из первого среза:", values_1)

	// //zadanie 3
	// data := []int{}
	// data_intput := 0
	// n := 1

	// for {
	// 	fmt.Print("Введите число: ")
	// 	fmt.Scan(&temp_input)
	// 	if temp_input == 0 {
	// 		break
	// 	}

	// data = append(data[:n], data[n+1:]...)
	// }
	// fmt.Println(data)

	// //zadanie 4
	// temp := []int{}
	// temp_input := 0
	// for {
	// 	fmt.Print("Введите температуру: ")
	// 	fmt.Scan(&temp_input)
	// 	if temp_input == 0 {
	// 		break
	// 	}

	// minTemp := slices.Min(temp)
	// fmt.Println("Минимальная температура:", minTemp, "°C")

	// maxTemp := slices.Max(temp)
	// fmt.Println("Максимальная температура:", maxTemp, "°C")

	// //zadanie 5
	// words := []string{}
	// words_input := ""

	// for {
	// 	fmt.Print("Введите строки: ")
	// 	fmt.Scan(&words_input)
	// 	if words_input == "Stop" {
	// 		break
	// 	}
	// 	words = append(words, words_input)
	//	fmt.Println(words)
	// }
	//slices.Reverse(words)
	// fmt.Println(words)

	//zadanie 6
	nums := []int{}
	nums_input := 0

	for {
		fmt.Print("Введите число: ")
		fmt.Scan(&nums_input)

		if nums_input == 0 {
			break
		}
		nums = append(nums, nums_input)
	}

	copy_nums := slices.Clone(nums) //копирую
	slices.Sort(copy_nums)          // сортирую
	fmt.Println(copy_nums)

	Sorted_1 := slices.Equal(nums, copy_nums) //сравниваю
	fmt.Println(Sorted_1)

	//zadanie 7
	myWishList := [3]string{"App", "Car", "PC-club"}
	friendsWishList := []string{"Apartment", "Car", "GOS"}
	registrationList := []string{}

	for _, gift := range myWishList {
		registrationList = append(registrationList, gift)
	}

	for _, gift := range friendsWishList {
		registrationList = append(registrationList, gift)
	}
	// registrationList :=append(myWishList[:],friendsWishList...)

	fmt.Println(registrationList)

	//zadanie 8

	toyList := [3]string{"Car", "Doll", "Ball"}

	testToyList := toyList

	testToyList[1] = "Boat"

	fmt.Println("Оригинал:", toyList)
	fmt.Println("Копия:", testToyList)
}
