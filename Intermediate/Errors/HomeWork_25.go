package main

import (
	"errors"
	"fmt"
)

// zadani 2
type error interface {
	Error() string
}

func divide(num1, num2 float64) (float64, interface{}) {
	if num2 == 0 {
		return 0, "Ошибка: деление на ноль невозможно"
	}
	return num1 / num2, nil
}

// zadanie 3
func checkAge(age int) error {
	if age < 0 || age > 120 {
		return errors.New("указан некорректный возраст")
	}
	return nil
}

//	func checkAge(age int) (int, error) {
//		if age < 0 || age > 120 {
//			return 0, errors.New("указан некорректный возраст")
//		}
//		return age, nil
//	}

// zadanie 4
func validatePassword(MyStr string) (bool, error) {

	if len(MyStr) < 6 {
		return false, errors.New("пароль слишком короткий, не менее 5 символов")
	}
	return true, nil
}

// zadanie 5
type invalidRadiusError struct{}

func (e *invalidRadiusError) Error() string {
	return ("Радиус не может быть отрицательным")
}

// zadanie 6
func calculateCircleArea(r int) (float64, error) {
	if r < 0 {
		return 0, &invalidRadiusError{}
	}
	return 3.14 * float64(r) * float64(r), nil
}

// zadanie 7
func findUser(names []string, search string) (int, error) {
	for i, n := range names {
		if n == search {
			return i, nil
		}
	}
	return -1, errors.New("пользователь не найден")
}

func main() {
	//zdanie 1
	m := map[string]int{
		"potato": 499,
		"Sprite": 650,
		"Nan":    340,
	}
	val, ok := m["Sprite"]
	if ok == true {
		fmt.Println(val)
	} else {
		fmt.Println("Товар отсутсвует")
	}

	//zadanie 2
	res, err2 := divide(52, 2)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(":", res)
	}

	res, err2 = divide(52, 0)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(":", res)
	}

	//zadanie 3
	if err3 := checkAge(-4); err3 != nil { //возвращает ошибку, елси надо вернуть и возраст то if _,err3:= checkAge...
		fmt.Println("Ошибка:", err3)

	}

	//zadanie 4

	if _, err4 := validatePassword("zxcv"); err4 != nil {
		fmt.Println("Ошибка:", err4)
	} else {
		fmt.Println("Ваш пароль подходит")
	}
	if _, err4 := validatePassword("1234678zxcv"); err4 != nil {
		fmt.Println("Ошибка:", err4)
	} else {
		fmt.Println("Ваш пароль подходит")
	}

	//zadanie 5
	radiusErr := &invalidRadiusError{}
	fmt.Println(radiusErr.Error())

	//zadanie 6
	if area, err6 := calculateCircleArea(-9); err6 != nil {
		fmt.Println("Ошибка:", err6.Error())
	} else {
		fmt.Println("Площадь:", area)
	}

	//zadanie 7
	Subscribers := []string{
		"Ardak",
		"Danik",
		"Agryn",
	}
	if index, err7 := findUser(Subscribers, "Ardak"); err7 != nil {
		fmt.Println("Ошибка", err7)
	} else {
		fmt.Println("Индекс саба:", index+1)
	}
}
