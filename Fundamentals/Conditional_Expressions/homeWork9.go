package main

import (
	"fmt"
)

func main() {
	//zadanie 1
	fmt.Println("1:", 5 == 5)                  //1:true
	fmt.Println("2:", 10 != 3)                 //2:true
	fmt.Println("3:", 7 > 12)                  //3:false
	fmt.Println("4:", 15 < 20)                 //4:true
	fmt.Println("5:", 8 >= 8)                  //5:true
	fmt.Println("6:", 6 <= 4)                  //6:false
	fmt.Println("7:", (10 > 5) && (3 < 1))     //7:false
	fmt.Println("8:", (10 > 5) || (3 < 1))     //8:true
	fmt.Println("9:", (!(5 == 5)))             //9:false
	fmt.Println("10:", !(7 < 3))               //10: true
	fmt.Println("11:", true && false)          //11:false
	fmt.Println("12:", false || false)         //12:false
	fmt.Println("13:", (true || false))        //13:true
	fmt.Println("14:", (4+6 == 10) && (9 > 2)) //14:true
	fmt.Println("15:", (12/3 == 4) || (8 < 5)) //15:true

	//zadanie 2
	hasTicket := true
	age := 52
	canEnter := hasTicket == true && age >= 18
	fmt.Println(canEnter)

	//zadanie 3

	isLoggedIn := true
	isAdmin := true
	hasAccess := (isLoggedIn && isAdmin) || (isLoggedIn && !isAdmin)
	fmt.Println(hasAccess)

}
