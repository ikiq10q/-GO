package main

import (
	"fmt"
)

type User struct {
	login, favouriteFood string
	age                  int
}

func increment_age(user User) {
	user.age++
}

func main() {

	user1 := User{
		login:         "user1",
		favouriteFood: "pizza",
		age:           25,
	}

	// fmt.Println(user1)

	// var user2 User
	// // fmt.Println(user2)
	// user2 = User{"user2", "sushi", 30}

	// fmt.Println(user1.favouriteFood)
	// fmt.Println(user2.favouriteFood)
	// fmt.Println()

	// user1.favouriteFood = "Pasta"
	// fmt.Println(user1.favouriteFood)
	// fmt.Println(user2.favouriteFood)
	// fmt.Println()

	// var p_user1 *User
	// p_user1 = &user1
	// fmt.Println(p_user1)

	// p_user1.favouriteFood = "pasta"
	// fmt.Println(p_user1)
	// fmt.Println(user1)

	// increment_age(user1)
	// fmt.Println(user1)

	user_copy := user1
	user_copy.age += 10
	fmt.Println(user1)
	fmt.Println(user_copy)

	fmt.Println(user1 == user_copy)
}
