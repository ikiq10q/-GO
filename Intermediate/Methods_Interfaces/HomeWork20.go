package main

import (
	"fmt"
	"strings"
)

// zadanie 1
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

//zadanie 2

type Student struct {
	Name   string
	Grades []int
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}

	score := 0.0
	for _, grade := range s.Grades {
		score += float64(grade)
	}

	return score / float64(len(s.Grades))
}

type Person interface {
	Average() float64
}

// zadanie 3
type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Printf("Баланс %s: %.2f \n", b.Owner, amount)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 && b.Balance >= amount {
		b.Balance -= amount
		fmt.Printf("Баланс %s: %.2f \n", b.Owner, amount)
	} else {
		fmt.Printf("Недостаточно средств для снятия %.2f (Баланс: %.2f).\n", amount, b.Balance)
	}
}

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
}

// zadanie 4
type TextAnalyzer struct {
	Text string
}

func (t TextAnalyzer) CountWords() int {
	words := strings.Fields(t.Text)
	return len(words)
}
func (t TextAnalyzer) CountSentences() int {
	count := 0
	for _, text := range t.Text {
		if text == '.' || text == '!' || text == '?' {
			count++
		}
	}
	return count
}
func (t TextAnalyzer) CountVowels() int {
	vowels := "aeiouyAEIOUY"
	count := 0

	for _, char := range t.Text {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}

type Analyzer interface {
	CountWords() int
	CountSentences() int
	CountVowels() int
}

func main() {
	//zadanie 1
	rect1 := Rectangle{Width: 14.1, Height: 5.2}
	rect2 := Rectangle{Width: 3.2, Height: 6.7}

	shapes := []Shape{rect1, rect2}

	for i, shape := range shapes {
		fmt.Printf("Прямоугольник %d: площадь %.2f, периметр %.2f\n", i+1, shape.Area(), shape.Perimeter())
	}

	// zadanie 2 трех студентов с разными оценками
	s1 := Student{Name: "Ахмед", Grades: []int{4, 5, 5, 4, 3}}
	s2 := Student{Name: "Ахмад", Grades: []int{5, 5, 5, 4, 4}}
	s3 := Student{Name: "Ахат", Grades: []int{2, 4, 3, 5, 2}}

	boys := []Person{s1, s2, s3}

	for _, b := range boys {

		if boys, ok := b.(Student); ok {
			fmt.Printf("Студент %s: средний балл %.2f\n", boys.Name, boys.Average())
		}
	}

	//zadanie 3
	acc1 := &BankAccount{Owner: "Матвей", Balance: 999.0}
	acc2 := &BankAccount{Owner: "Пудж", Balance: 667.0}

	accounts := []Account{acc1, acc2}

	fmt.Printf("Баланс %s: %.2f тг \n", acc1.Owner, acc1.Balance)
	fmt.Printf("Баланс %s: %.2f тг \n", acc2.Owner, acc2.Balance)

	acc1.Deposit(500)
	acc1.Withdraw(200)

	acc2.Withdraw(1000)
	acc2.Deposit(100)

	for _, acc := range accounts {

		if bAcc, ok := acc.(*BankAccount); ok {
			fmt.Printf("Баланс %s: %.2f тг \n", bAcc.Owner, bAcc.Balance)
		}
	}

	//zadanie 4

	analyzer := TextAnalyzer{
		Text: "Легендарная фраза: \"Alhumdulillah, I'm gonna smash your boy!\", в UFC",
	}

	fmt.Printf(
		"Слов: %d, предложений: %d, гласных: %d\n",
		analyzer.CountWords(),
		analyzer.CountSentences(),
		analyzer.CountVowels(),
	)

}
