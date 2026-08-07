package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books map[string]Book

// ищем книгу по id
func bookHandler(c *gin.Context) {
	id := c.Param("id")

	book, ok := books[id]
	if !ok {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "Книга с ID %s не найдена", id)
		return
	}

	fmt.Fprintf(c.Writer, "Книга: %s, Автор: %s", book.Title, book.Author)
}

// обновляем книгу
func updateBookHandler(c *gin.Context) {
	id := c.Param("id")

	book, ok := books[id]
	if !ok {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "Книга с ID %s не найдена", id)
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books[id] = book
	c.JSON(http.StatusOK, book)
}

// zadanie 2
var users map[string]User

type User struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func deleteUserHandler(c *gin.Context) {
	username := c.Param("username")

	// Проверка на админа
	if username == "admin" {
		c.Status(http.StatusForbidden)
		fmt.Fprintf(c.Writer, "Ошибка: нельзя удалить пользователя %s", username)
		return
	}

	// Проверка пользователя
	_, ok := users[username]
	if !ok {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "Пользователь %s не найден", username)
		return
	}

	delete(users, username)

	c.Status(http.StatusOK)
	fmt.Fprintf(c.Writer, "Пользователь %s успешно удален", username)
}

// zadanie 5
type ServerStatus struct {
	Status        string `json:"status"`
	UptimeSeconds int    `json:"uptime_seconds"`
}

func statusHandler(c *gin.Context) {
	var nowStatus ServerStatus

	fmt.Print("Введите статус:")
	fmt.Scan(&nowStatus.Status)
	nowStatus.Status = strings.ToLower(nowStatus.Status)

	fmt.Print("Введите время работы: ")
	fmt.Scan(&nowStatus.UptimeSeconds)

	err := c.ShouldBindJSON(&nowStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nowStatus)

}
func main() {
	//zadanie 1
	books = map[string]Book{
		"1": {ID: "1", Title: "123", Author: "Karm"},
		"2": {ID: "2", Title: "322", Author: "DDD"},
		"3": {ID: "3", Title: "4444", Author: "XXD"},
	}
	//zadanie 2
	users = map[string]User{
		"admin":     {Username: "Ardak", Role: "Admin"},
		"creepchek": {Username: "Kamik", Role: "User"},
	}
	r := gin.Default()
	r.GET("/books/:id", bookHandler)
	r.GET("/books/:id", updateBookHandler)

	//zadanie 2
	r.DELETE("/users/:username", deleteUserHandler)

	r.GET("/status", statusHandler) //zadanie 5

	fmt.Println("Сервер запущен на 8080 порту")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}

}
