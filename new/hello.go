package new

import (
	"fmt"
)

func SayHello() {
	fmt.Println("WW")

	var pf *float64
	fmt.Println("Value: ", *pf) // ! ошибка, указатель не указывает на какой-либо объект

}
