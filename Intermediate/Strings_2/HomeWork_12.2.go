package main

import (
	"fmt"
	"strings"
	"unicode"
)

// zadanie 7
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// zadanie 8
func hasThreeConsecutive(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return true
		}
	}
	return false
}

// zadanie 9
func checkPassword(s string) {
	if len(s) < 8 {
		fmt.Println("Слишком короткий пароль")
		return
	}
	hasDig, hasUp := false, false
	for _, pass := range s {
		if unicode.IsDigit(pass) {
			hasDig = true
		}
		if unicode.IsUpper(pass) {
			hasUp = true
		}
	}
	if !hasDig {
		fmt.Println("Пароль должен содержать цифру")
	} else if !hasUp {
		fmt.Println("Пароль должен содержать заглавную букву")
	} else {
		fmt.Println("Пароль корректный")
	}
}

// zadanie 10
func checkEmail(email string) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		fmt.Println("Ошибка: Email должен содержать один @")
		return
	}
	if !strings.Contains(parts[1], ".") {
		fmt.Println("Ошибка:Email должен содержать точку после @")
		return
	}
	fmt.Println("Email корректен")
}

// zadanie 11
func checkFormat(s string) {
	if len(s) == 0 {
		return
	}

	frstRune := []rune(s)[0]
	starts := unicode.IsUpper(frstRune)
	ends := strings.HasSuffix(s, ".")

	if !starts {
		fmt.Println("Не с заглавной буквы")
	}
	if !ends {
		fmt.Println("Не заканчивается точкой")
	}
	if starts && ends {
		fmt.Println("Строка оформлена правильно")
	}
}

// zadanie 12
func countLetters(s string) {
	aCount, eCount := 0, 0
	for _, r := range s {
		if r == 'a' || r == 'A' {
			aCount++
		}
		if r == 'e' || r == 'E' {
			eCount++
		}
	}
	fmt.Printf("a/A: %d, e/E: %d\n", aCount, eCount)
}
func main() {
	//zadanie 1
	fmt.Println(strings.ToUpper("GoLang"))
	fmt.Println(strings.ToLower("GoLang"))

	//zadanie 2
	fmt.Println(strings.TrimSpace(" backend developer "))

	//zadanie 3
	str := "I like Java"
	if strings.Contains(str, "Java") {
		fmt.Println("Найдено")
	} else {
		fmt.Println("Не найдено")
	}

	//zadanie 4
	Z4 := "one,two,three"
	res := strings.ReplaceAll(Z4, ",", ";")
	fmt.Println(res)

	//zadanie 5
	fmt.Print(" zadanie 5 Введите данные: ")
	var input string
	fmt.Scanln(&input)
	if strings.Index(input, "admin") == 0 {
		fmt.Println("Доступ разрешён")
	} else {
		fmt.Println("Доступ запрещён")
	}

	//zadanie 6
	s := "Go is fast and Go is simple"
	count := strings.Count(s, "Go")
	fmt.Println(count)

	//zadanie 7
	fmt.Print(" zadanie 7 Введите строку: ")
	var z7 string
	fmt.Scanln(&z7)
	if isPalindrome(z7) {
		fmt.Println("Это палиндром")
	} else {
		fmt.Println("Это не палиндром")
	}

	//zadanie 8
	fmt.Print("zadanie 8 Введите строку: ")
	var z8 string
	fmt.Scanln(&z8)
	fmt.Println(hasThreeConsecutive(z8))

	//zadanie 9
	fmt.Print("zadanie 9 Введите пароль: ")
	var z9 string
	fmt.Scanln(&z9)
	checkPassword(z9)

	//zadanie 10
	fmt.Print("zadanie 10 Введите email: ")
	var z10 string
	fmt.Scanln(&z10)
	checkEmail(z10)

	//zadanie 11
	fmt.Print("zadanie 11 Введите предложение: ")
	var z11 string
	fmt.Scanln(&z11)
	checkFormat(z11)

	//zadanie 12
	fmt.Print("zadanie 12 Введите строку: ")
	var z12 string
	fmt.Scanln(&z12)
	countLetters(z12)
}
