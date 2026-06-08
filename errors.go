package main

import (
	"errors"
	"fmt"
)

// type error interface {
// 	Error() string
// }

// type zero_devision struct{}

// func (err zero_devision) Error() string {
// 	return "Нeльзя делить на 0"
// }

// type grater_than_100000 struct{}

// func (err grater_than_100000) Error() string {
// 	return "Я не хочу считать такое большое число"
// }

func dev(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Нeльзя делить на 0")
	}
	if a > 100000 {
		return 0, errors.New("Я не хочу считать такое большое число")
	}
	return a / b, nil
}

func main1() {
	a, ok := dev(10, 0)

	if ok == nil {
		fmt.Println(a)
	} else {
		fmt.Println(ok)
	}

	fmt.Println(dev(10, 0))
	if ok == nil {
		fmt.Println(a)
	} else {
		fmt.Println(ok.Error())
	}
	fmt.Println(dev(100000000000000, 5))
	if ok == nil {
		fmt.Println(a)
	} else {
		fmt.Println(ok.Error())
	}

}
