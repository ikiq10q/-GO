package main

import "fmt"

func main() {
	str := "Hello"
	rune_str := []rune(str)
	rune_str[1] = 'a'

	str = string(rune_str)

	fmt.Println(str)
}
