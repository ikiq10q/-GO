package main

import (
	"fmt"
	// "sort"
)

func main1() {
	// initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

	// users1 := initialUsers[2:6] // с 3-го по 6-й
	// users2 := initialUsers[:4]  // с 1-го по 4-й
	// users3 := initialUsers[3:]  // с 4-го до конца

	// fmt.Println(users1)
	// fmt.Println(users2)
	// fmt.Println(users3)

	// var users []int = make([]int, 3) // длина среза - 3

	// users[0] = 9
	// users[1] = 4
	// users[2] = 5

	// for i, v := range users {
	// 	fmt.Printf("Index: %d, value: %d\n", i, v)
	// }

	user_2 := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	clients := make([]string, len(user_2)*2)

	copy(clients, user_2)
	fmt.Println(clients)
	clients[6] = "Argyn"
	fmt.Println(clients)
}
