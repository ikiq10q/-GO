package main

import (
	"fmt"
	"strings"
)

func main() {
	//zadanie 1
	text := "Hi dude"
	const author = "Ardak"

	fmt.Printf("Автор %s написал: %s.\n", author, text)
	long := len(text)
	fmt.Printf("Длина строки: %d символов\n", long)

	//zadanie 2
	message := "Легендарная фраза: \"Alhumdulillah, I'm gonna smash your boy!\", в UFC"

	if message == "" {
		fmt.Println("Строка пуста")
	} else {
		fmt.Println("Строка не пуста")
	}
	messageLong := len(message)
	fmt.Println("Длина сообщения:", messageLong)

	//zadanie 3
	fmt.Print("Введите строку: ")

	input := ""
	fmt.Scan(&input)
	length := len(input)

	if length < 5 {
		fmt.Println("Слишком короткое слово")
	} else if length >= 5 && length <= 10 {
		fmt.Println("Нормальная длина")
	} else {
		fmt.Printf("Слишком длинное слово")
	}

	//zadanie 4

	word := "Ardak"

	firstSymb := string(word[0])
	lastSymb := string(word[len(word)-1])

	fmt.Println("Первый символ:", firstSymb)
	fmt.Println("Последний символ:", lastSymb)

	fmt.Println("Длина строки:", len(word))

	fmt.Println("Все символы строки:")

	for _, symb := range word {
		fmt.Println(string(symb))
	}

	//zadanie 5
	sentence := "Задание по строкам в го"

	runes := []rune(sentence)

	total := len(runes)

	countA := 0

	for _, v := range runes {

		if v == 'а' || v == 'А' {
			countA++
		}
	}

	fmt.Printf("В строке %d символов и %d букв а\n", total, countA)

	//zadanie 6
	fmt.Print("Введите строку: ")

	scanner := ""
	fmt.Scan(&scanner)
	var reversed string

	for i := len(scanner) - 1; i >= 0; i-- {
		reversed += string(scanner[i])

	}
	fmt.Println("Введеная строка:", scanner)
	fmt.Println("Перевернутая строка:", reversed)

	if scanner == reversed {
		fmt.Println("Это палиндром")
	} else {
		fmt.Println("Это не палиндром")
	}

	//zadanie 7
	fmt.Print("Введите строку: ")

	text2 := ""
	fmt.Scan(&text2)

	length = len(text2) // длина строки
	words := strings.Fields(text2)
	wordCount := len(words)

	vowels := "aeiouyAEIOUY" // считаем гласные
	vowelCount := 0

	for _, v := range text2 {
		if strings.ContainsRune(vowels, v) {
			vowelCount++
		}
	}

	fmt.Printf("Длина строки: %d, слов: %d, гласных: %d\n",
		length, wordCount, vowelCount)

}
