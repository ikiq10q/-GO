package main

import (
	"fmt"
)

type User struct {
	name         string
	login        string
	age          int
	isAdmin      bool
	userNoteBook NoteBook
}

type NoteBook struct {
	nCores   int
	size     float64
	keyBoard KeyBoard
}

type KeyBoard struct {
	nkeys int
}

func (u User) print() {
	fmt.Println("Name:", u.name)
	fmt.Println("Login:", u.login)
	fmt.Println("Age:", u.age)
	fmt.Println("Admin?:", u.isAdmin)
	fmt.Println("NoteBook:", u.userNoteBook)
}

func NewUser(name string, login string, age int) *User {
	var nCores int
	var userNote float64
	fmt.Println("Введите колличество ядер и размер")
	fmt.Scan(&nCores, &userNote)

	return &User{name, login, age, false, NoteBook{nCores, userNote, KeyBoard{120}}}

}

func (u *User) setAdmin() {
	u.isAdmin = true

}

func (u *User) setLogin(newLogin string) {
	if u.name == "Oleg" {
		u.login = newLogin

	}
}

func (u *User) getLogin() string {
	if u.name == "Oleg" {
		return u.login
	}
	return "У вас нет прав на редактирование"
}
func main() {
	person := NewUser("Oleg", "Oleggs", 20)
	person.print()

	person.setAdmin()

	person.setLogin("Oleggwp")
	person.print()
	fmt.Println(person.getLogin())
}
