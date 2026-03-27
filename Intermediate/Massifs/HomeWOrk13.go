package main

import "fmt"

func main() {
	//zadanie 1
	running := [4]string{"Кросс", "Спринт", "Берпи", "Растяжка"}
	walking := [3]string{"Спортивная хольба", "Приседания", "Растяжка"}

	fmt.Println(len(running))
	fmt.Println(len(walking))

	//zadanie 2
	subjectsList := [3]string{"Физика", "Химия", "География"}

	first := subjectsList[0]
	last := subjectsList[len(subjectsList)-1]
	subjectsList[1] = "Биология"

	fmt.Println(first)
	fmt.Println(last)
	fmt.Println(subjectsList)

	//zadanie 3
	character := [3]string{"Tom", "35", "New York"}

	name := character[0]
	age := character[1]
	city := character[2]

	fmt.Printf("%v\n%v\n%v\n", name, age, city)

	//zadanie 4
	numberList := [5]int{1, 2, 3, 4, 5}

	for _, element := range numberList {
		if element == 3 {
			fmt.Println("Число 3 есть в массиве")
		} else {
			fmt.Println("Число 3 отсутствует в массиве")
		}
	}

	//zadanie 5
	freidsList := [3]string{"Agyn", "Alexei", "Danik"}
	for _, name := range freidsList {
		if name == "Bekbolat" {
			fmt.Println("Мне очень повезло")
			return
		}
	}

	fmt.Println("Мне не повезло")

	//zadanie 6
	firstList := [3]int{1, 2, 3}
	secondList := [3]int{1, 2, 4}

	if firstList == secondList {
		fmt.Println("Массивы равны")
	} else {
		fmt.Println("Массивы не равны")
	}

	//zadanie 7
	myWishList := [2]string{"Mouse", "Car"}
	friendsWishList := [2]string{"Car", "Painting"}

	var registrationList []string

	for _, gift := range myWishList {
		registrationList = append(registrationList, gift)
	}

	for _, gift := range friendsWishList {
		registrationList = append(registrationList, gift)
	}

	fmt.Println(registrationList)

	//zadanie 8
	toyList := [3]string{"Car", "Doll", "Ball"}
	testToyList := toyList

	testToyList[1] = "Boat"

	fmt.Println(toyList)
	fmt.Println(testToyList)

}
