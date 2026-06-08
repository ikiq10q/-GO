package main

import (
	"fmt"
)

func end() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func dev(a, b float64) float64 {
	defer end()

	if b == 0 {
		panic("Dev by 0")
	}
	return a / b
}

func main() {
	a := 10
	if a := dev(120, 60); a != 0 {
		fmt.Println(a)
	}
	fmt.Println(a)
	fmt.Println("Program ended")
	fmt.Println("Program ended")
	fmt.Println("Program ended")
}
