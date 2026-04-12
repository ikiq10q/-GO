package main

import "fmt"

//Zadanie 1
func main() {
	age := 28
	fmt.Println("Age:", age)
	age += 1
	fmt.Println("Age:", age)

	//Zadanie 2
	height := 170
	fmt.Println("Height:", height, "sm")

	height_in_meters := 1.7
	fmt.Println("Height_in_meters:", height_in_meters, "m")

	//Zadanie 3
	isStudent := true
	fmt.Println(isStudent)

	//Zadanie 4
	temperature := -25
	fmt.Println("Temperature:", temperature, "°C")
	fmt.Println("Погода холодная")

	//Zadanie 5
	favoriteQuote := "Alhumdulillah, I'm gonna smash your boy!"
	fmt.Println("My favorite Quote", favoriteQuote)

	//Zadanie 6
	PI := 3.14
	fmt.Println("PI:", PI)

	// PI = "3.1415"
	// fmt.Println("PI:", PI)
	// не получится присвоить "PI" строку так как она уже имееет значение float64,
	// так же если было бы const
}
