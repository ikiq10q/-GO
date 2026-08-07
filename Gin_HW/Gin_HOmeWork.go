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

// zadanie 3
var products map[string]Product

type Product struct {
	Title    string `json:"title"`
	PriceKZT int    `json:"price_KZT"`
}

func productHandler(c *gin.Context) {
	sku := c.Param("sku")

	p, ok := products[sku]
	if !ok {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "Товар с артикулом %s отсутствует", sku)
		return
	}
	priceUSD := p.PriceKZT / 450

	c.Status(http.StatusOK)
	fmt.Fprintf(c.Writer, "Товар: %s, Цена: %d KZT (примерно %d USD)", p.Title, p.PriceKZT, priceUSD)
}

// zadanie 4
var orders map[string]Order

type Order struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func OrderStatusDelivered(c *gin.Context) {
	id := c.Param("id")
	newStatus := c.Param("new_status")

	//проверка на наличии заказа
	order, ok := orders[id]
	if !ok {
		c.Status(http.StatusNotFound)
		fmt.Fprintf(c.Writer, "Заказ %s не найден", id)
		return
	}
	//статсус заказа
	validStatuses := map[string]bool{
		"pending":   true,
		"shipped":   true,
		"delivered": true,
	}
	// проверка на валидацию
	if !validStatuses[newStatus] {
		c.Status(http.StatusBadRequest)
		fmt.Fprintf(c.Writer, "Статус %s недопустим", newStatus)
		return
	}

	order.Status = newStatus
	orders[id] = order // сохраняем обратно

	c.Status(http.StatusOK)
	fmt.Fprintf(c.Writer, "Статус заказа %s успешно изменен на %s", id, newStatus)

}

// zadanie 4.1
type WeatherReport struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
}

func weatherHandler(c *gin.Context) {
	var report WeatherReport

	err := c.ShouldBindJSON(&report)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
	fmt.Fprintf(c.Writer, "Город: %s, Температура: %.1f°C", report.City, report.Temperature)
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

	//zadanie 3
	products = map[string]Product{
		"sku001": {Title: "Хлеб", PriceKZT: 123},
		"sku002": {Title: "Cy", PriceKZT: 321},
		"sku003": {Title: "Груши", PriceKZT: 322},
	}

	//zadanie 4
	orders = map[string]Order{
		"001": {ID: "001", Status: "pending"},
		"002": {ID: "002", Status: "delivered"},
	}

	r := gin.Default()
	r.GET("/books/:id", bookHandler)
	r.GET("/books/:id", updateBookHandler)

	//zadanie 2
	r.DELETE("/users/:username", deleteUserHandler)

	//zadanie 3
	r.GET("/products/:sku", productHandler)

	r.POST("/orders/:id/status/:new_status", OrderStatusDelivered) //zadanie 4

	r.POST("/weather", weatherHandler) //zadanie 4.1

	r.GET("/status", statusHandler) //zadanie 5

	fmt.Println("Сервер запущен на 8080 порту")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}

}
